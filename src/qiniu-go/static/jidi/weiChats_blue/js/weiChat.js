$(function(){
    //点击非弹窗主要内容区域，关闭弹窗
	$(".popup").mousedown(function(e){
		var target = $(e.target);
        //点击其他地方隐藏弹窗
        if(!target.is('.popup-box *')&&!target.is('.popup-box')) {
           $(".popup").hide();
        }
	});
	//点击关闭弹窗按钮
	$(".popup-box-close").click(function(){
		$(this).parents(".popup").hide();
	});
	$(".close").click(function(){
		$("ul.redpacket_list-ul").scrollTop(0);
		$(this).parents(".bigPopup").hide();
	});
	//点击游客--关注
	$(document).on("click",".popup-box-flow_btn",function(){
		$(this).addClass("popup-box-flow_focused").text("已关注");
	});
	
	//点击抢红包	
	$(".visitor_sendRedpacket_box").click(function(){
		$(".grab_redpacket").show();
	});
	$(".now_redpacket").click(function(){
		$(".grab_redpacket").show();
	});
	//红包详情
	$(document).on("click",".grab_redpacket-box-details",function(){
		$(this).parents(".popup").hide();
		$(".redpacket_list").addClass("show");
		$(".redpacket_list").show();
	});
	
	$("ul.redpacket_list-ul").scroll(function(){
		$(".redpacket_list-theLastli").slideUp(300);
		if($(this).scrollTop()==0){
			$(".redpacket_list-theLastli").slideDown(500);
		}
	});
	//点击礼物按钮
	$(".chat-box_gif-btnBox").click(function(){
		$(".send_gift").show();
	});
	
	//点击礼物主要内容区域，关闭弹窗
	$(".send_gift").mousedown(function(e){
		var target = $(e.target);
        //点击其他地方隐藏弹窗
        if(!target.is('.send_gift-box *')&&!target.is('.send_gift-box')) {
           $(".send_gift").hide();
        }
	});
	$('.send_gift').on('click', 'ul.send_gift-ul li', function(){
		//选中高亮
		$(".send_gift ul.send_gift-ul li").removeClass("active").find(".send_gift-ul-money").hide();
		$(this).addClass("active");
		$(this).find(".send_gift-ul-money").show();
		//获取当前礼物信息
		giftActiveInfor = $("ul.send_gift-ul li.active");
		giftid=giftActiveInfor.attr("data-giftid");
		giftname=giftActiveInfor.attr("data-giftname");
		giftimg=giftActiveInfor.attr("data-giftimg");
		gitfmoney=giftActiveInfor.attr("data-giftmoney");
		giftcolor=giftActiveInfor.attr("data-giftcolor");
		//获取用户总余额
		remainderNumber=parseInt($(".send_gift-remainder i").text());
	});
    
    //点击游客头像
	$(document).on("click",".other_words img",function(){
		$(".user_infor").show();
	});
    $(document).on("click",".seqencingHeader-img  img",function(){
		$(".user_infor").show();
	});
    
	//贡献榜切换
	$("ul.seqencing-tab li").click(function(){
		$("ul.seqencing-tab li").removeClass("active");
		$(this).addClass("active");
		var liIndex=$(this).index();
		$(".seqencingBox>div").hide().eq(liIndex).show();
	});
	//贡献榜--关注
	$(".seqencingBox-ul li>span.seqencingUser-indexOrfollow").click(function(){
		if($(this).hasClass("seqencingUser-index")){
		}else{
			$(this).css("background","none").text("已关注")
		}
	});
	//二维码
	$(".fixed_fun ul li.fixed_fun-erweima").click(function(){
		$(".erweima").show();
	});
	//发红包
	$(".fixed_fun-redpacket").click(function(){
		$(".send_redpacket").addClass("show");
		$(".send_redpacket").show();
	});
	//打赏
	$(".fixed_fun-award").click(function(){
		$(".award_anchor").addClass("show");
		$(".award_anchor").show();
		
	});
	//打赏星钻 --输入框
	var input_awardMoneyval;
	$(".input_awardMoney").keyup(function(e){
		//判断输入金额是否合法----合法启用btn_awardMoney按钮
		input_awardMoneyval=parseInt($("input.input_awardMoney").val().trim());
		var reg = /[^\d]/g;
		var input_awardMoneyvalReg=reg.test(input_awardMoneyval);
		$(".award_anchor ul.award_anchor_ul li span").removeClass("active");
		if(!input_awardMoneyvalReg && input_awardMoneyval>0){
			$(".btn_awardMoney").removeClass("btn_disable").addClass("btn_able").removeAttr("disabled");
			console.log(input_awardMoneyval)
		}
		else{
			$(".btn_awardMoney").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
			$(".errorMoney_point").text("打赏不低于一颗星钻");
			//错误提示
			var errorMoney_point_width=$(".errorMoney_point").width();
			$(".errorMoney_point").css("left","50%").css("margin-left",-(errorMoney_point_width/2 + 5)+"px");
				if(e.keyCode!=8){
					errorMoney_point();
					console.log(input_awardMoneyval)
			}
		}
	});
	//选择打赏金额
		$(".award_anchor ul.award_anchor_ul li span").click(function(){
			$(".award_anchor ul.award_anchor_ul li span").removeClass("active");
			$(this).addClass("active");
			var awardvalue = $(this).find("input").val()
			$(".input_awardMoney").val(awardvalue);
			$(".btn_awardMoney").removeClass("btn_disable").addClass("btn_able").removeAttr("disabled");
		});

	//点击赏
	/**
	 * 1.获取用户余额
	 * 2.比较打赏金额和用户余额
	 * 3.打赏成功清空数据，禁用btn_awardMoney按钮,关闭弹窗
	 */
	$(".btn_awardMoney").click(function(){
		$(".award_anchor ul.award_anchor_ul li span").removeClass("active");
		remainderNumber=parseInt($(".send_gift-remainder>i").text());
		input_awardMoneyval=parseInt($("input.input_awardMoney").val().trim());
		if(remainderNumber<input_awardMoneyval){
			$(".notEnough_popup").show();
		}else if(remainderNumber>=input_awardMoneyval){
			parseInt($(".send_gift-remainder>i").text(remainderNumber-input_awardMoneyval));
			$("input.input_awardMoney").val("");
			$(this).removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
			$(this).parents(".popup").fadeOut();
		}
	});
	
	
	//输入红包个数
	var redpacket_num;
	$(".redpacket_num").keyup(function(e){
		redpacket_num=$(".redpacket_num").val().trim();
		if(redpacket_num<1){
			$(".btn_send_redpacket").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
			$(".errorMoney_point").text("红包个数不低于1个");
			if(e.keyCode!=8){
				errorMoney_point();
			}
		}else if(redpacket_all/redpacket_num<0.01 || redpacket_num%1!=0){
			$(".btn_send_redpacket").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
		}else if( redpacket_all<=0 || redpacket_all==undefined){
			$(".btn_send_redpacket").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
		}else{
			$(".btn_send_redpacket").removeClass("btn_disable").addClass("btn_able").removeAttr("disabled");
		}
	});
	
	
	//输入红包总金额
	var redpacket_all;
	$(".redpacket_all").keyup(function(e){
		redpacket_all=$(".redpacket_all").val().trim();
		if(redpacket_all<=0 || redpacket_all/redpacket_num<0.01){
			$(".btn_send_redpacket").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
			$(".errorMoney_point").text("单个红包不低于0.01元");
			if(e.keyCode!=8){
				errorMoney_point();
			}
		}else if( redpacket_num<1 || redpacket_num%1!=0 || redpacket_num==undefined) {
			$(".btn_send_redpacket").removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
		}else{
			$(".btn_send_redpacket").removeClass("btn_disable").addClass("btn_able").removeAttr("disabled");
		}
	});
	
	
	
	//点击塞钱
	$(".btn_send_redpacket").click(function(){
		redpacket_all=0;
		redpacket_num=0;
		$("input.redpacket_num").val("");
		$("input.redpacket_all").val("");
		$("input.btn_send_redpacket").val("");
		$(this).removeClass("btn_able").addClass("btn_disable").attr("disabled","disabled");
		$(this).parents(".popup").fadeOut();
	});
	
	
	
	
	//点击分享
	$(".fixed_fun-share").click(function(){
		$(".share_alert").show();
	});
	$(".share_alert").click(function(){
		$(this).hide();
	})
	
	//点击非分享区域关闭分享
	$(".popup-share").mousedown(function(e){
		var target = $(e.target);
        //点击其他地方隐藏弹窗
        if(!target.is('.share-box *')&&!target.is('.share-box')) {
           $(".popup-share").hide();
        }
	});
	
	//点击友情链接
	$(".friendshipLink").click(function(){
		if($(".friendshipLink_popup").css("display")=="none"){
			var friendshipLink_boxLength = $(".friendshipLink_div>div");
				if(friendshipLink_boxLength.length<=1){
					$(".friendshipLink_div").css("width","120px");
					$(".friendshipLink_box").css("width","100%");
					$(".friendshipLink_box").css("margin-bottom","0");
				}else if(friendshipLink_boxLength.length<3&&friendshipLink_boxLength.length>1){
					$(".friendshipLink_div").css("width","240px");
					$(".friendshipLink_box").css("width","120px");
					$(".friendshipLink_box").css("margin-bottom","0");
				}
			
			$(".friendshipLink_fullPop").show();
			$(".friendshipLink_popup").fadeIn(); 
		}
	});
	
	//点击非友情链接区域关闭友情链接
	$(".friendshipLink_fullPop").mousedown(function(e){
		var target = $(e.target);
        //点击其他地方隐藏弹窗
        if(!target.is('.friendshipLink_popup *')&&!target.is('.friendshipLink_popup')) {
        	$(".friendshipLink_fullPop").hide();
           $(".friendshipLink_popup").hide();
        }
	});
	
	//点击关注主播
	$(".focus_heart").mousedown(function(){
		if($(this).hasClass("heart")){
			$(this).find("img").attr("src","img/heartx2.png");
			$(this).removeClass("heart");
			$(".focus_heartInfor").text("已关注");
		}else{
			$(this).find("img").attr("src","img/heartx1.png");
			$(this).addClass("heart");
			$(".focus_heartInfor").text("关注");
		}
	});
	
	//点击门票详情
	$(".detailAndbuy a.detail-a").click(function(){
		$(".orderDetail").show();
	});
	
	//////购买弹窗/////
	$(".close-paymentDetails").click(function(){
		$(".paymentDetails_popup").hide()
	})
		
	$(".notEnough_popup .otherWayPay").click(function(){
		$(".notEnough_popup").hide();
		$(".paymentDetails_popup").addClass("show").show();
	});
	
	$(".chosePay-box2").click(function(){
		$(".payment-box").hide();
		$(".way-box").show();
	});
	$(".goback-paymentBox").click(function(){
		$(".way-box").hide();
		$(".payment-box").show();
	});
	$(".way-box label").click(function(){
		var choiceImg = $(this).find(".choiceImg img").attr("src");
		var choicevalue=$(this).find(".choiceThisWay input").val()
		console.log(choicevalue)
		$("label.chosePay-box2 .chosePay-box-img img").attr("src",choiceImg);
		$("label.chosePay-box2 input.chosePay-box2-radio").val(choicevalue);
	})
	
//	$(".chosePay-box>label.chosePay-box-label").click(function(){
//		var chosePayIndex = $(this).index();
//		if(chosePayIndex==0){
//			$(this).find("img").attr("src","img/weichart.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(1).find("img").attr("src","img/balance2x.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(2).find("img").attr("src","img/star2x.png");
//		}
//		if(chosePayIndex==1){
//			$(this).find("img").attr("src","img/balance.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(0).find("img").attr("src","img/weichart2x.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(2).find("img").attr("src","img/star2x.png");
//		}
//		if(chosePayIndex==2){
//			$(this).find("img").attr("src","img/star.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(0).find("img").attr("src","img/weichart2x.png");
//			$(".chosePay-box>label.chosePay-box-label").eq(1).find("img").attr("src","img/balance2x.png");
//		}
//	});
	
//	$(".chosePay-box input").click(function(){
//		alert($(this).val())
//	})
	
});//$over

	//获取聊天列表 -- 使滚动条保持在底部
	function holdBottom(){
		chatList = document.getElementsByClassName('chat-list')[0];
		chatList.scrollTop=chatList.scrollHeight;
	}
	//打赏or红包金额过少提示弹框动画
	function errorMoney_point(){
		$(".errorMoney_point").show();
		$(".errorMoney_point").animate({top:'4%', opacity:'1'},800,function(){
			$(".errorMoney_point").animate({top:'4%', opacity:'1'},800,function(){
				$(".errorMoney_point").animate({top:0, opacity:0},100).css("display","none");
			});
		});
	}
	