namespace go base

struct User{
    1: string id;
    2: string username;
    3: string mobile;
    4: string gender;
    5: string age;
}

struct UserInfo{
    1: string id;
    2: string username;
    3: string mobile;
    4: string gender;
    5: string age;
    6: string avatar_url;
}




//日志信息
struct LogsInfo {
    1:  string type (api.raw = "type")
    2:  string method (api.raw = "method")
    3:  string api (api.raw = "api")
    4:  string success (api.raw = "success")
    5:  string reqContent (api.raw = "reqContent")
    6:  string respContent (api.raw = "respContent")
    7:  string ip (api.raw = "ip")
    8:  string userAgent (api.raw = "userAgent")
    9:  string operator (api.raw = "operator")
    10:  string time (api.raw = "time")
    11:  string createdAt (api.raw = "createdAt")
    12:  string updatedAt (api.raw = "updatedAt")
}


// 字典信息
struct DictionaryInfo {
    1:  string id (api.raw = "id" )
    2:  string title (api.raw = "title" )
    3:  string name (api.raw = "name" )
    5:  string status (api.raw = "status" )
    6:  string description (api.raw = "description" )
    7:  string createdAt (api.raw = "createdAt" )
    8:  string updatedAt (api.raw = "updatedAt" )
}
// 字典键值信息
struct DictionaryDetail {
    1:  string id (api.raw = "id" )
    2:  string title (api.raw = "title" )
    3:  string key (api.raw = "key" )
    4:  string value (api.raw = "value" )
    5:  string status (api.raw = "status" )
    6:  string createdAt (api.raw = "createdAt" )
    7:  string updatedAt (api.raw = "updatedAt" )
    8:  string parentID (api.raw = "parentID" )
}