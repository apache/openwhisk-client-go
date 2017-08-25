/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package whisk

import (
    "io/ioutil"
    "fmt"
    "net/url"
    "os"
    "strings"
    "bufio"
    "errors"
    "github.com/apache/incubator-openwhisk-client-go/wski18n"
)

const (
    OPENWHISK_HOME = "OPENWHISK_HOME"
    GOPATH = "GOPATH"
    DEFAULT_LOCAL_CONFIG = ".wskprops"
    OPENWHISK_PROPERTIES = "whisk.properties"
    TEST_AUTH_FILE = "testing.auth"
    OPENWHISK_PRO= "whisk.api.host.proto"
    OPENWHISK_PORT= "whisk.api.host.port"
    OPENWHISK_HOST= "whisk.api.host.name"
    DEFAULT_VERSION = "v1"
    DEFAULT_NAMESPACE = "guest"
    DEFAULT_API = "api"

    NAMESPACE = "NAMESPACE"
    AUTH = "AUTH"
    WHISK_APIHOST = "WHISK_APIHOST"
    APIGW_ACCESS_TOKEN = "APIGW_ACCESS_TOKEN"
    APIVERSION = "APIVERSION"
    KEY = "KEY"
    CERT = "CERT"
    APIHOST = "APIHOST"

    DEFAULT_SOURCE = ".wskprops"
    WSKPROP = ".wskprops"
    WHISK_PROPERTY = "whisk.properties"
)

type Wskprops struct {
    WHISKAPIURL string
    APIHost string
    AuthKey string
    Namespace string
    AuthAPIGWKey string
    APIGWSpaceSuid string
    Apiversion string
    Key string
    Cert string
    Source string
}

func convertWskpropsToConfig(dep *Wskprops) *Config {
    var config Config
    var err error
    config.BaseURL, err = url.Parse(dep.WHISKAPIURL)
    if err != nil {
        config.BaseURL = nil
    }
    config.Namespace = dep.Namespace
    config.Cert = dep.Cert
    config.Key = dep.Key
    config.AuthToken = dep.AuthKey

    if config.BaseURL != nil {
        config.Host = config.BaseURL.Host
    } else {
        config.Host = ""
    }

    config.Version = dep.Apiversion
    config.Verbose = false
    config.Debug = false
    config.Insecure = true

    return &config
}

func GetDefaultConfigFromProperties(pi Properties) (*Config, error) {
    var config *Config
    dep, e := GetDefaultWskProp(pi)
    config = convertWskpropsToConfig(dep)
    return config, e
}

func GetConfigFromWhiskProperties(pi Properties) (*Config, error) {
    var config *Config
    dep, e := GetWskPropFromWhiskProperty(pi)
    config = convertWskpropsToConfig(dep)
    return config, e
}

func GetConfigFromWskprops(pi Properties, path string) (*Config, error) {
    var config *Config
    dep, e := GetWskPropFromWskprops(pi, path)
    config = convertWskpropsToConfig(dep)
    return config, e
}

var GetDefaultWskProp = func (pi Properties) (*Wskprops, error) {
    var dep *Wskprops
    dep = pi.GetPropsFromWskprops("")
    error := ValidateWskprops(dep)
    if error != nil {
        dep_whisk := pi.GetPropsFromWhiskProperties()
        error_whisk := ValidateWskprops(dep_whisk)
        if error_whisk != nil {
            return dep, error
        } else {
            return dep_whisk, error_whisk
        }
    }
    return dep, error
}

var GetWskPropFromWskprops = func (pi Properties, path string) (*Wskprops, error) {
    var dep *Wskprops
    dep = pi.GetPropsFromWskprops(path)
    error := ValidateWskprops(dep)
    return dep, error
}

var GetWskPropFromWhiskProperty = func (pi Properties) (*Wskprops, error) {
    var dep *Wskprops
    dep = pi.GetPropsFromWhiskProperties()
    error := ValidateWskprops(dep)
    return dep, error
}

type Properties interface {
    GetPropsFromWskprops(string) *Wskprops
    GetPropsFromWhiskProperties() *Wskprops
}

type PropertiesImp struct {
    OsPackage OSPackage
}

func (pi PropertiesImp) GetPropsFromWskprops(path string) *Wskprops {
    dep := GetDefaultWskprops(WSKPROP)

    var configPath string
    if path != "" {
        configPath = path
    } else {
        configPath = pi.OsPackage.Getenv(GOPATH, "")
    }
    results, err := ReadProps(configPath + "/" + DEFAULT_LOCAL_CONFIG)

    if err == nil {

        dep.WHISKAPIURL = GetValue(results, WHISK_APIHOST, dep.WHISKAPIURL)
        baseURL, err := url.Parse(dep.WHISKAPIURL)
        if err != nil {
            dep.APIHost = GetValue(results, APIHOST, dep.APIHost)
        } else {
            dep.APIHost = baseURL.Host
        }
        dep.AuthKey = GetValue(results, AUTH, dep.AuthKey)
        dep.Namespace = GetValue(results, NAMESPACE, dep.Namespace)
        dep.AuthAPIGWKey = GetValue(results, APIGW_ACCESS_TOKEN, dep.AuthAPIGWKey)
        if len(dep.AuthKey) > 0 {
            dep.APIGWSpaceSuid = strings.Split(dep.AuthKey, ":")[0]
        }
        dep.Apiversion = GetValue(results, APIVERSION, dep.Apiversion)
        dep.Key = GetValue(results, KEY, dep.Key)
        dep.Cert = GetValue(results, CERT, dep.Cert)
    }

    return dep
}

func (pi PropertiesImp) GetPropsFromWhiskProperties() *Wskprops {
    dep := GetDefaultWskprops(WHISK_PROPERTY)
    path := pi.OsPackage.Getenv(OPENWHISK_HOME, "") + "/" + OPENWHISK_PROPERTIES
    results, err := ReadProps(path)

    if err == nil {
        authPath := GetValue(results, TEST_AUTH_FILE, "")
        b, err := ioutil.ReadFile(authPath)
        if err == nil {
            dep.AuthKey = strings.TrimSpace(string(b))
        }

        var pro = GetValue(results, OPENWHISK_PRO, "")
        var port = GetValue(results, OPENWHISK_PORT, "")
        var host = GetValue(results, OPENWHISK_HOST, "")
        dep.WHISKAPIURL = fmt.Sprintf("%s://%s:%s/%s", pro, host, port, DEFAULT_API)
        baseURL, err := url.Parse(dep.WHISKAPIURL)
        if err != nil {
            dep.APIHost = fmt.Sprintf("%s:%s", host, port)
        } else {
            dep.APIHost = baseURL.Host
        }
        dep.Namespace = DEFAULT_NAMESPACE
        if len(dep.AuthKey) > 0 {
            dep.APIGWSpaceSuid = strings.Split(dep.AuthKey, ":")[0]
        }
    }
    return dep
}

var ValidateWskprops = func (wskprops *Wskprops) error {
    // There are at least two fields: WHISKAPIURL and AuthKey, mandatory for a valid Wskprops.
    errStr := ""
    if (len(wskprops.WHISKAPIURL) == 0) {
        if wskprops.Source == WHISK_PROPERTY {
            errStr = wski18n.T("OpenWhisk API URL is missing (Please configure WHISK_APIHOST in .wskprops under the system HOME directory.)")
        } else {
            errStr = wski18n.T("OpenWhisk API URL is missing (Please configure whisk.api.host.proto, whisk.api.host.name and whisk.api.host.port in whisk.properties under the OPENWHISK_HOME directory.)")
        }
        return MakeWskError(errors.New(errStr), EXIT_CODE_ERR_GENERAL, DISPLAY_MSG, DISPLAY_USAGE)
    } else {
        _, err := url.ParseRequestURI(wskprops.WHISKAPIURL)
        if (err != nil) {
            if wskprops.Source == WHISK_PROPERTY {
                errStr = wski18n.T("Invalid OpenWhisk API URL (Please configure WHISK_APIHOST in .wskprops under the system HOME directory.)")
            } else {
                errStr = wski18n.T("Invalid OpenWhisk API URL (Please configure whisk.api.host.proto, whisk.api.host.name and whisk.api.host.port in whisk.properties under the OPENWHISK_HOME directory.)")
            }
            return MakeWskError(errors.New(errStr), EXIT_CODE_ERR_GENERAL, DISPLAY_MSG, DISPLAY_USAGE)
        } else if (len(wskprops.AuthKey) == 0) {
            if wskprops.Source == WHISK_PROPERTY {
                errStr = wski18n.T("Authentication key is missing (Please configure AUTH in .wskprops under the system HOME directory.)")
            } else {
                errStr = wski18n.T("Authentication key is missing (Please configure testing.auth as the path of the authentication key file in whisk.properties under the OPENWHISK_HOME directory.)")
            }
            return MakeWskError(errors.New(errStr), EXIT_CODE_ERR_GENERAL, DISPLAY_MSG, DISPLAY_USAGE)
        } else {
            return nil
        }
    }
}

type OSPackage interface {
    Getenv(key string, defaultValue string) string
}

type OSPackageImp struct {}

func (osPackage OSPackageImp) Getenv(key string, defaultValue string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return defaultValue
    }
    return value
}

func GetDefaultConfig() (*Config, error) {
    pi := PropertiesImp{
        OsPackage: OSPackageImp{},
    }
    return GetDefaultConfigFromProperties(pi)
}

func GetWhiskPropertiesConfig() (*Config, error) {
    pi := PropertiesImp{
        OsPackage: OSPackageImp{},
    }
    return GetConfigFromWhiskProperties(pi)
}

func GetPropertiesImp() PropertiesImp {
    return PropertiesImp{
        OsPackage: OSPackageImp{},
    }
}
func GetWskpropsConfig(path string) (*Config, error) {
    pi := GetPropertiesImp()
    return GetConfigFromWskprops(pi, path)
}

func GetDefaultWskprops(source string) *Wskprops {
    if len(source) == 0 {
        source = DEFAULT_SOURCE
    }

    dep := Wskprops {
        WHISKAPIURL: "",
        APIHost: "",
        AuthKey: "",
        Namespace: DEFAULT_NAMESPACE,
        AuthAPIGWKey: "",
        APIGWSpaceSuid: "",
        Apiversion: DEFAULT_VERSION,
        Key: "",
        Cert: "",
        Source: source,
    }
    return &dep
}

func GetValue(StoredValues map[string]string, key string, defaultvalue string) string {
    if val, ok := StoredValues[key]; ok {
        return val
    } else {
        return defaultvalue
    }
}

func ReadProps(path string) (map[string]string, error) {

    props := map[string]string{}

    file, err := os.Open(path)
    if err != nil {
        return props, err
    }
    defer file.Close()

    lines := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    props = map[string]string{}
    for _, line := range lines {
        kv := strings.Split(line, "=")
        if len(kv) != 2 {
            continue
        }
        key := strings.TrimSpace(kv[0])
        value := strings.TrimSpace(kv[1])
        props[key] = value
    }

    return props, nil

}
