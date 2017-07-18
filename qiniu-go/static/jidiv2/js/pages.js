	//////////////////////////////////////全局///////////////////////////
	//关闭弹窗
	$(document).on("click",".close-pagespopup",function(){
		$(".popup-box").html("");
        $(".popup").hide();
	});
	$(document).on("click",".popupCloseButton-min",function(){
		$(".popup-box").html("");
        $(".popup").hide();
	})
	//关闭弹窗
	$(".popup").click(function(e){
		var target = $(e.target);
        if(!target.is('.popup-box>div *')&&!target.is('.popup-box>div')) {
        	$(".popup-box").html("");
        	$(".popup").hide();
        }
	});
	//////////////////////////////////////全局----///////////////////////////
	
	////////////////////////////////收益////////////////////////
	$(".incomewarp .topChoice .choice").click(function(){
		let choiceIndex = $(this).index();
		$(".incomewarp ul.box").eq(choiceIndex).show().siblings("ul.box").hide();
	});
	
	////////////////////////////////我的消费////////////////////////
	
	//礼物
	//礼物详情
	$(document).on("click",".salesOrder-giftBox-show",function(){
		$(".popup").show();
        var data={};
        var html = template('giftsDetails-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	//门票
	//门票详情
	$(document).on("click",".salesOrder-ticketBox-show",function(){
		$(".popup").show();
        var data={};
        var html = template('ticketDetails-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	//商品
	//显示二维码
	$(document).on("click",".salesOrder-contact-show",function(e){
		$(".qr-code_popup").show()
	})
	
	//隐藏二维码
	$(".qr-code_popup").click(function(e){
		var target = $(e.target);
		if(!target.is('.code_popupBox')&&!target.is('.code_popupBox *')){
			$(".copy-button").text("复制")
			$(".qr-code_popup").hide();
		}
	})
//	$(document).on("click",".qr-code_popup",function(e){
//		var target = $(e.target);
//		if(!target.is('.code_popupBox')&&!target.is('.code_popupBox *')){
//			$(".copy-button").text("复制")
//			$(".qr-code_popup").hide();
//		}
//	});
	
	//处于可交易的商品
	$(document).on("click",".salesOrder-goodsBox-inTheDeal-show",function(){
		$(".popup").show();
        var data={};
        var html = template('salesOrder-goodsBox-inTheDeal-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	 
	//处于退款的商品
	function _refundmentStatus(){
    	 //退款状态--等待卖家核实 status-wait  退款中  status-refunding 退款成功 status-success
	     var returnStatus="status-success";
	     var progressBar = $(".salesOrder-popup .number-pb .number-pb-shown");
	     var progressCircle = $(".salesOrder-popup .number-pb .number-pb-num");
	     var progressLi = $(".salesOrder-popup .number-pb .number-ul li")
	     if(returnStatus==="status-wait"){
	     	progressBar.css("width","33.333333%");
	     	progressCircle.css("left","33.333333%");
	     	progressLi.eq(0).css("color","#499989");
	     }else if(returnStatus==="status-refunding"){
	     	progressBar.css("width","66.666666%");
	     	progressCircle.css("left","66.666666%");
	     	progressLi.eq(1).css("color","#499989");
		 }else if(returnStatus==="status-success"){
		 	progressBar.css("width","100%");
	     	progressCircle.css("left","calc(100% - 10px)");
	     	progressLi.eq(2).css("color","#499989");
		 }
    }
	$(document).on("click",".salesOrder-goodsBox-refunding-show",function(){
		$(".popup").show();
        var data={};
        var html = template('salesOrder-goodsBox-returnPolicy-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
        _refundmentStatus();
	});
