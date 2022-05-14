package templates

const GetOrderById = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>L0 Product info</title>
    <style>
        .myDiv {
            border: 2px outset black;
            width: fit-content;
            padding: 10px;
            background: whitesmoke;
            text-align: center;
        }
    </style>
</head>
<body>
<h3>Input order Id</h3>
<form method="POST">
    <label>Order Id</label><br>
    <input type="text" name="id" /><br><br>
    <input type="submit" value="Get info" /><br>
</form><br>
<label>RESULT</label><br>
<div class="myDiv">
    <p>{{ .}}</p>
</div>
</body>
</html>`
