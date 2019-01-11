<!DOCTYPE html>

<html>
<head>
  <title>BeegoDemo</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">  
</head>

<body>
  <div>
	<form action="/query" method="Post">
		<div>
			GetPaymentBy AccountID:<input  type="text" name="AccountID1" />
		</div>
				
		<div>
			<input type= "submit" name="n" />
		</div>
	</form>
	<form action="/add" method="Post">
	<div>
			 AccountID:<input  type="text" name="AccountID" />
	</div>
	<div>
		 PartnerID:<input  type="text" name="PartnerID" />
	</div>
	<div>
		 UserID   :<input  type="text" name="UserID" />
	</div>
	<div>
		 CreateTime:<input  type="text" name="CreateTime" />
	</div>
	<div>
		 Amount:<input  type="text" name="Amount" />
	</div>
	<div>
		 OuterTradeNo:<input  type="text" name="OuterTradeNo" />
	</div>
	<div>
		 Remark:<input  type="text" name="Remark" />
	</div>

	<div>
		<input type= "submit" name="add" value="增加"/>
	</div>
</form>
	
  </div>
</body>
</html>
