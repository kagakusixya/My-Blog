<!DOCTYPE HTML>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <title>
        KG掲示板
    </title>
</head>
<body>
<h1>KG掲示板:{{.key1}}{{.key2}}</h1>
<h3><pre></pre></h3>
<hr>
<form action="/Formcontrollers" method="post">
    名前:<input type="text" name="name"><br>
    コメント:<input type="text" name="say"><br>
    <input type="submit" value="送信">
    {{.key1}}:{{.key2}}
</form>
</body>
</html>
