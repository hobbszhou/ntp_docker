# ntp服务



## 1.获取ntp服务的系统时间及时区



**功能描述**

**Url:  http://10.51.15.193:8085/api/v1/ntp/ntpInfo**

**请求方式:  GET**

**返回参数:**

| 返回字段 | 二级字段  | 字段类型 | 说明                                             |
| -------- | --------- | -------- | :----------------------------------------------- |
| 说明     |           | int      | 返回结果状态。200:正常; 400:错误;500:错误， 其它 |
| msg      |           | string   | ok;其它                                          |
| result   | time_zone | 数组     | ntp服务端的所有时区                              |
|          | ntp_time  | string   | ntp服务的时间                                    |

**返回示例:**

```json

成功示例:
{
    "code": 200,
    "message": "ok",
    "result": {
        "time_zone": [
            "Africa/Abidjan",
            "Africa/Accra",
            "Africa/Addis_Ababa",
            "Africa/Algiers",
            "Africa/Asmara",
            "Africa/Bamako",
            "Africa/Bangui",
            "Africa/Banjul",
            "Africa/Bissau",
            "Africa/Blantyre",
            "Africa/Brazzaville",
            "Africa/Bujumbura",
            "Africa/Cairo",
            "Africa/Casablanca",
            "Africa/Ceuta",
            "Africa/Conakry",
            "Africa/Dakar",
            "Africa/Dar_es_Salaam",
            "Africa/Djibouti",
            "Africa/Douala",
            "Africa/El_Aaiun",
            "Africa/Freetown",
            "Africa/Gaborone",
            "Africa/Harare",
            "Africa/Johannesburg",
            "Africa/Juba",
            "Africa/Kampala",
            "Africa/Khartoum",
            "Africa/Kigali",
            "Africa/Kinshasa",
            "Africa/Lagos",
            "Africa/Libreville",
            "Africa/Lome",
            "Africa/Luanda",
            "Africa/Lubumbashi",
            "Africa/Lusaka",
            "Africa/Malabo",
            "Africa/Maputo",
            "Africa/Maseru",
            "Africa/Mbabane",
            "Africa/Mogadishu",
            "Africa/Monrovia",
            "Africa/Nairobi",
        ],
        "ntp_time": "Wed Jul 14 14:29:19 CST 2021"
    }
}

```

## 2.设置ntp服务

**功能描述:** 设置ntp服务时区及指定ntp服务远端节点

**Url:  http://10.51.15.193:8085/api/v1/ntp/ntpServer**

**请求方式:  post**

**请求参数:** 

| 请求字段        | 类型   | 是否必填 | 说明       |
| :-------------- | ------ | -------- | ---------- |
| time_zone       | string | 是       | 时区       |
| ntp_remote_addr | string | 是       | 远端ip地址 |

**请求示例：**

```json

{
        "time_zone":"Asia/Shanghai",
        "ntp_remote_addr":"10.51.15.194"
}


```

**返回参数:**



| 返回字段 | 二级字段 | 字段类型 | 说明                                             |
| -------- | -------- | -------- | ------------------------------------------------ |
| code     |          | int      | 返回结果状态。200:正常; 400:错误;500:错误， 其它 |
| msg      |          | string   | ok;其它                                          |
| result   |          | string   |                                                  |

**返回示例**

```
成功:
{
    "code": 200,
    "message": "ok",
    "result": "set success"
}
```



































