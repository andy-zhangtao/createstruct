# createstruct
The gojsonschema rest api

## How to use?

目前只支持Json Draft 4所规定的规范。 为了方便上传json schema，在此约定schema必须进行base64编码

如果转换失败，则也会返回一个空struct。所以无法通过返回码来判断转换是否成功，只能通过返回的字符串来判断转换是否成功

## API LIST

API VERSION :/v1

<hr/>

Method|Path|Description|Parameter|
------|----|-----------|---------|
GET|/_ping|测试后端服务是否正常|无|
POST|/post/json|使用上传的schema生成struct数据|Body: <br/> {name:'必填','json':'必填'}|