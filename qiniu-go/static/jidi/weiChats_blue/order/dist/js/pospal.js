$(".navbtn li a").click(function(e){
	e.preventDefault();
	$(this).parent().addClass("active").siblings().removeClass("active");
	$($(this).attr("href")).addClass("blo").siblings().removeClass("blo");
});

//阅读并同意
$(function(){
	$(".foot p span").click(function(){
		var val = $(this).find("img").attr("val");
		if(val==0){
			$(this).find("img").attr({"src":"dist/img/icon-no.png","val":"1"});
		}
		if(val==1){
			$(this).find("img").attr({"src":"dist/img/icon-yes.png","val":"0"});
		}
	});
});

var wait=60;
function time(o) { 
		
        if (wait == 0) { 
            o.removeAttribute("disabled");           
            o.value="获取验证码"; 
            wait = 60; 
        } else { 
            o.setAttribute("disabled", true); 
            o.value="重新发送("+wait+")"; 
            wait--; 
            setTimeout(function() { 
                time(o) 
            }, 
            1000) 
        } 
    } 
var waita=60;
function timea(o) { 
		
        if (waita == 0) { 
            o.removeAttribute("disabled");           
            o.value="获取验证码"; 
            waita = 60; 
        } else { 
            o.setAttribute("disabled", true); 
            o.value="重新发送("+waita+")"; 
            waita--; 
            setTimeout(function() { 
                timea(o) 
            }, 
            1000) 
        } 
    } 
   
$(".starbtn-proving").click(function(){
	time(this);
});
$(".cashbtn-proving").click(function(){
	timea(this);
}); 
$(".shade-pnoneproving").click(function(){
	timea(this);
}); 

$(".off").click(function(){
	$(this).parents(".shade").hide();
}); 
$(".changecon").click(function(){
	$(this).next(".shade").show();
}); 

//var data1 = [
//			{
//				value: 12345,
//				color:"#229ff7"
//			},
//			{
//				value : 2000,
//				color : "#fe4b4b"
//			}			
//		]
//		new Chart(document.getElementById("starbalance").getContext("2d")).Pie(data1);


