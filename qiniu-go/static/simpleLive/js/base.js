	//关闭弹窗
	$(".popup").click(function(e){
		var target = $(e.target);
        if(!target.is('.popup-box>div *')&&!target.is('.popup-box>div')) {
        	$(".collection-classify-popup").removeClass("show");
        	$(".popup-box").html("");
        	$(".popup").hide();
        }
	});
	//关闭弹窗
	$(".popup-right").click(function(e){
		var target = $(e.target);
        if(!target.is('.popup-box-right>div *')&&!target.is('.popup-box-right>div')) {
        	$(".collection-classify-popup").addClass("hide");
        	setTimeout(function(){
        		$(".popup-box-right").html("");
        		$(".popup-right").hide();
        	},300)
        }
	});
	$(document).on("click",".closebtn-popup",function(){
		$(".popup-box").html("");
        $(".popup").hide();
	});
	$(document).on("click",".popupCloseButton-min",function(){
		$(".popup-box").html("");
        $(".popup").hide();
	})
	//顶部切换
	$(".custom-topmenu li a").click(function(){
		$(this).addClass("active").parent("li").siblings("li").find("a").removeClass("active");
	});
	
	//左边一级列表切换样式
	$(document).on("click",".sidebar-menu>li.up-menu",function(){
		$(this).addClass("active").siblings("li").removeClass("active");
	});
	
	//右边一级内容---显示隐藏
	$(document).on("click",".rightcont",function(event){
		event.stopPropagation();
		//加载层-风格3
		layer.load();
		//此处演示关闭
		setTimeout(function(){
		  layer.closeAll('loading');
		}, 500);
		var dataRightBoxCont = $(this).attr("data-rightBox");
		console.log("右边内容" + dataRightBoxCont);
		$("#main-content section.wrapper-" + dataRightBoxCont).addClass("active").siblings(".wrapper").removeClass("active");
		
		if(dataRightBoxCont=="FinancialCenter"){
			var testBox = $(".sidebar-menu li[data-rightbox=FinancialCenter]");
			testBox.addClass("active").siblings("li").removeClass("active");
		}
		if(dataRightBoxCont=="PersonalInformation"){
			var testBox = $(".sidebar-menu li[data-rightbox=PersonalInformation]");
			testBox.addClass("active").siblings("li").removeClass("active");
		}
	});
	//左边二级列表切换样式
	$(document).on("click",".sidebar-menu>li.sub-menu>ul.sub>li",function(){
		$(this).addClass("active").siblings("li").removeClass("active");
	});
	//右边二级内容---显示隐藏
	$(document).on("click",".custom-nav-tabs li",function(event){
		event.stopPropagation();
		var navtabsliIndex = $(this).index();
		$(this).parent("ul.custom-nav-tabs").siblings(".wrapper-content").eq(navtabsliIndex).addClass("active").siblings().removeClass("active")
	});
	
	
	//收货地址/回放管理/门票--删除提示
	$(document).on("click",".operation-popup-show",function(){
		$(".popup").show();
        var data;
        var deleteListObjType = $(this).attr("data-trashtype");
        //删除回放
		if(deleteListObjType==="VideoBox"){
			data={
				title:"确定要删除该回放吗？"
			}
		}
		//删除收货地址
		if(deleteListObjType==="AddressBox"){
			data={
				title:"确定要删除该收货地址吗？"
			}
			var trList = $("table.address-tbl tbody tr");
			if(trList.length<2){
				$(".wrapper-content-ShippingAddress .address-data-box").hide();
			}
		}
		//删除门票
		if(deleteListObjType==="TicketBox"){
			data={
				title:"确定要删除该门票吗？"
			}
		}
		//删除友情链接
		if(deleteListObjType==="FriendsOfOheOhainBox"){
			data={
				title:"确定要删除该友情链接吗？"
			}
		}
		//删除房管
		if(deleteListObjType==="RoomHousingManagementBox"){
			data={
				title:"确定要取消该房管吗？"
			}
		}
		
		//取消拉黑
		if(deleteListObjType==="RoomBlacklistBox"){
			data={
				title:"确定要取消拉黑吗？"
			}
		}
		//取消禁言
		if(deleteListObjType==="RoomGaglistBox"){
			data={
				title:"确定要取消禁言吗？"
			}
		}
		//删除简介
		if(deleteListObjType==="RomeSynopsisBox"){
			data={
				title:"确定要删除该简介吗？"
			}
		}
       	var html = template('operation-popup', data);
        $('.popup-box').html(html);
        var deleteBtn = $(this);
        _deleteListFun(deleteListObjType,deleteBtn)
	});
	function _deleteListFun(deleteListObjType,deleteBtn){
			$('.delete-yes').off('click').on("click",function(){
				deleteBtn.parents(".operation-trashObj").remove();
				$(".popup-box").html("");
       			$(".popup").hide();
			});
	}
	
	
	//拉黑
	$(document).on("click",".pullToBlack-popup-show",function(e){
		var blackBtn = $(this);
		$(".popup").show();
        var data={};
        var html = template('pullToBlack-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
        $('.pullToBlack-yes').off('click').on("click",function(){
			blackBtn.parent("td.table-details").html('<span data-trashtype="RoomBlacklistBox" class="table-details operation-popup-show">取消拉黑</span>')
			$(".popup-box").html("");
   			$(".popup").hide();
		});
	});
	
	
	
	///////////////财务中心///////////////////////
	//账单查看详情
	$(document).on("click",".bill-details-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('bill-details-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	//星票提现
	$(document).on("click",".startickets-withdraw-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('startickets-withdraw-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	//现金提现
	$(document).on("click",".cash-withdraw-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('cash-withdraw-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	
	///////////////我的消费///////////////////////
	//门票详情
	$(document).on("click",".ticketDetails-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('ticketDetails-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	//商品--交易中弹窗
	$(document).on("click",".salesOrder-goodsBox-inTheDeal-popup-show",function(){
		$(".popup").show();
		 var data={};
		var salesOrderType = $(this).attr("data-salesOrder")
		if(salesOrderType==="GoodsToBeReceived"){
			data.btnType="GoodsToBeReceived"
		}
       	else if(salesOrderType==="SuccessfulTrade"){
			data.btnType="SuccessfulTrade"
		}
       	else if(salesOrderType==="ShipmentPending"){
			data.btnType="ShipmentPending"
		}
       	console.log(data.btnType)
        var html = template('salesOrder-goodsBox-inTheDeal-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	$(document).on("click",".inTheDeal_button button.refundRequest",function(){
		$(".popup-box").html("");
		var data={};
        var html = template('salesOrder-goodsBox-returnPolicy-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
        var returnPolicy_button = '<div class="returnPolicy_button"><button class="inline-block closebtn-popup" type="button">提交申请</button></div>';
        $(".salesOrder-goodsBox-returnPolicy-popup.salesOrder-popup").append(returnPolicy_button);
	});
	
	
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
	});
	
	//处于退款的商品
	function _refundmentStatus(){
    	 //退款状态--等待卖家核实 status-wait  退款中  status-refunding 退款成功 status-success
	     var returnStatus="status-wait";
	     var progressBar = $(".salesOrder-popup .number-pb .number-pb-shown");
	     var progressCircle = $(".salesOrder-popup .number-pb .number-pb-num");
	     var progressLi = $(".salesOrder-popup .number-pb .number-ul li")
	     if(returnStatus==="status-wait"){
	     	progressBar.css("width","33.333333%");
	     	progressCircle.css("left","33.333333%");
	     	progressLi.eq(1).css("color","#499989");
	     }else if(returnStatus==="status-refunding"){
	     	progressBar.css("width","66.666666%");
	     	progressCircle.css("left","66.666666%");
	     	progressLi.eq(2).css("color","#499989");
		 }else if(returnStatus==="status-success"){
		 	progressBar.css("width","100%");
	     	progressCircle.css("left","calc(100% - 10px)");
	     	progressLi.eq(3).css("color","#499989");
		 }
    }
	$(document).on("click",".salesOrder-goodsBox-refunding-show",function(){
        $(".popup").show();
		var data={
			refundingbtnType:0
		};
		var salesOrderType = $(this).attr("data-salesOrder")
		
        var html = template('salesOrder-goodsBox-returnPolicy-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
        
        if(salesOrderType==="Refunding"){
			data.refundingbtnType="Refunding";
			$("input.return-policy-input").attr("disabled","disabled");
			$("textarea.return-policy-textarea").attr("disabled","disabled");
		}
       	else if(salesOrderType==="Successrefunding"){
			data.refundingbtnType="Successrefunding";
			$("input.return-policy-input").attr("disabled","disabled");
			$("textarea.return-policy-textarea").attr("disabled","disabled");
		}
       	console.log(data.refundingbtnType)
        
        _refundmentStatus();
	});
	
	//礼物详情
	$(document).on("click",".gifts-details-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('gifts-details-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	///////////////个人信息////////////////////////////////
	function indicatorContainer(){
		$('#indicatorContainer').radialIndicator({
		    barColor: '#446BA6',
		    radius: 35,//圆大小
		    initValue : 40,//页面加载是不动态加载数据，固定在某个值
		    percentage: true,//百分比
		    displayNumber: false,//中间是否显示数据
		    barWidth: 3//精度条粗细
		});
		var radialObj = $('#indicatorContainer').data('radialIndicator');
		radialObj.animate(40);
	}
	
	//修改姓名
	$(".person_infor-modifyName").click(function(){
		$(this).hide();
		$(this).siblings("span").hide();
		$(this).siblings("p").show();
	});
	$(document).click(function(e){
		var target = $(e.target);
        if(!target.is('.form-modifyname>input')&&!target.is('.person_infor-modifyName')) {
        	$(".form-modifyname").hide();
        	$(".form-modifyname").siblings("span").show();
			$(".form-modifyname").siblings("a").show();
        }
	});
	//头像设置
	$(".tab-box_userHeader").hover(function(){
		$(".change_userHeader").show();
	},function(){
		$(".change_userHeader").hide();
	});
	
	//绑定手机号
	$(document).on("click",".privacysecurity-phone-pupup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('privacysecurity-phone-pupup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	//绑定邮箱
	$(document).on("click",".privacysecurity-mailbox-pupup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('privacysecurity-mailbox-pupup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	//修改密码
	$(document).on("click",".privacysecurity-password-pupup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('privacysecurity-password-pupup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	
	
	
	///////////////////////视频管理///////////////////////////////////
	//回放列表切换
	$(document).on("click",".wrapper-content-playback-tab>ul.playback-tab-ul-a>li",function(){
		$(this).addClass("active").siblings("li").removeClass("active");
		var VideoManagementTab = $(this).attr("data-VideoManagementTab");
		if(VideoManagementTab=="all"){
			$(".backVideo-btn-group>.backVideo-tab-btn-a").eq(1).show().siblings("span").hide();
		}else if(VideoManagementTab=="up"){
			$(".backVideo-btn-group>.backVideo-tab-btn-a").eq(0).show().siblings("span").hide();
		}else{
			$(".backVideo-btn-group>span").hide();
		}
	});
	
	
	//视频设置
	$(document).on("change",".videoEditor-popup .form-group-videoEditor-selectTicket",function(){
		if($(this).val()==="videoTicketNotConfigured"){
			$(".videoTicketPasswordSet-input").removeAttr("disabled","disabled")
			$(".videoTicketPasswordSet-input").attr("placeholder","")
		}else{
			$(".videoTicketPasswordSet-input").attr("disabled","disabled")
			$(".videoTicketPasswordSet-input").attr("placeholder","门票和视频加密不可同时设置")
		}
	});
	
	//编辑视频
	$(document).on("click",".videoEditor-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('videoEditor-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	
	//上传视频
	$(document).on("click",".videoUpload-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('videoUpload-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	//分享
	$(document).on("click",".share-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('share-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});	
	
	//合并回放视频
	$(document).on("click",".backVideo-tab-btn-a button.combine-btn-b",function(){
		var videoCheckbox = '<input type="checkbox" class="wrapper-content-table-vBox-checkbox" />'
		var videoBoxObj = $(".wrapper-content-table .wrapper-content-table-vBox")
		if(videoBoxObj){
			videoBoxObj.find(".video").append(videoCheckbox)
			$(".wrapper-content-table-vBox-checkbox").show();
			$(".backVideo-tab-btn-c").show();
			$(".backVideo-tab-btn-b").hide();
			$(".backVideo-tab-btn-a").hide();
			$(".playback-tab-ul-a").hide();
		}else{
			return false;
		}
	});
	
	//合并
	$(document).on("click",".backVideo-tab-btn-c button.combine-btn-c",function(){
		var checkboxLen = $(".wrapper-content-table-vBox-checkbox:checked")
		if(checkboxLen.length>0){
			
			
		}else{
			alert("请选择合并的视频")
			return false;
		}
	});
	
	//完成
	$(document).on("click",".backVideo-tab-btn-c button.combine-btn-d",function(){
			$(".backVideo-tab-btn-a").show();
			$(".backVideo-tab-btn-c").hide();
			$(".playback-tab-ul-a").show();
			$(".video .wrapper-content-table-vBox-checkbox").remove();
	});
	
	
	//添加到标签
	$(document).on("click",".collection-classify-popup-show",function(){
		$(".popup-right").show();
        var data={};
        var html = template('collection-classify-popup', data);
        document.getElementsByClassName('popup-box-right')[0].innerHTML = html;
        $(".collection-classify-popup").addClass("show");
	});
	//标签设置
	$(document).on("click",".collection-classify-popup .collection-classify-ul li",function(){
		$(this).toggleClass("active")
	});
	
	
	
	//////////////////标签管理////////////////////
	//添加新标签
	$(document).on("click",".VideoLabel-edit-addNew-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('VideoLabel-edit-addNew-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	//标签编辑
	$(document).on("click",".VideoLabel-manage-popup-show",function(){
		$(".popup").show();
        var data={};
        var html = template('VideoLabel-manage-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	//删除标签中的视频
	$(document).on("click",".VideoLabel-manage-popup .trash-VideoLabel-manage-list",function(){
		var liObj = $(this);
		liObj.parents("li").animate({right:"999px"},500,function(){
			$(this).remove()
		});
//		liObj.parents("li").fadeOut(300, function(){
//			$(this).remove()
//		})
//		setTimeout((function(){
//			console.log(liObj)
//			$(".VideoLabel-manage-popup .VideoLabel-manage-ul li.remove").remove();
//		})(liObj),300)
		
//		(function(){
//			console.log(liObj)
//			$(".VideoLabel-manage-popup .VideoLabel-manage-ul li.remove").remove();
//		})(liObj)
		
	})
	
	////////////////////门票及商品/////////////////////////////	
	
	//上架、下架
	$(document).on("click",".GoodsShelves-popup-show",function(){
		var lowerORupper = $(this).attr("data-GoodsShelves")
		$(".popup").show();
        var data={
        	lowerORupper:lowerORupper
        };
        var html = template('GoodsShelves-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	
	//商户信息编辑
	
	$(document).on("click",".BusinessInformationEdit-popup-show",function(){
		$(".popup").show();
		var data={};
        var html = template('BusinessInformationEdit-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	//商城商品编辑
	$(document).on("click",".mall-CommodityEditor-popup-show",function(){
		var commodityAddOrEdit = $(this).attr("data-commodityAddOrEdit");
		var data;
		if(commodityAddOrEdit==="add"){
			data={
				title:"添加商品"
			}
		}else if(commodityAddOrEdit==="edit"){
			data={
				title:"编辑商品"
			}
		}
		$(".popup").show();
        var html = template('mall-CommodityEditor-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	
	//添加门票、编辑门票
	$(document).on("click",".EntranceTicket-edit-popup-show",function(){
		var ticketAddOrEdit = $(this).attr("data-ticketAddOrEdit");
		var data;
		if(ticketAddOrEdit==="add"){
			data={
				title:"添加门票"
			}
		}else if(ticketAddOrEdit==="edit"){
			data={
				title:"编辑门票"
			}
		}
		$(".popup").show();
        var html = template('EntranceTicket-edit-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	//修改分类
	$(document).on("click",".wrapper-content-RoomClassification .RoomClassification-classify-ul li",function(){
		$(this).addClass("active").siblings("li").removeClass("active")
	});
	
	//友链编辑
	$(document).on("click",".FriendlyLink-editor-popup-show",function(){
		var friendlyLinkAddOrEdit = $(this).attr("data-friendlyLinkAddOrEdit");
		var data;
		if(friendlyLinkAddOrEdit==="add"){
			data={
				title:"添加友链"
			}
		}else if(friendlyLinkAddOrEdit==="edit"){
			data={
				title:"友链编辑"
			}
		}
		$(".popup").show();
        var html = template('FriendlyLink-editor-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
	});
	
	
	
	
	//简介编辑
	$(document).on("click",".layui-edit-popup-show",function(){
		
		var RomeSynopsisAddOrEdit = $(this).attr("data-RomeSynopsisAddOrEdit");
		var data;
		if(RomeSynopsisAddOrEdit==="add"){
			data={
				title:"添加简介"
			}
		}else if(RomeSynopsisAddOrEdit==="edit"){
			data={
				title:"编辑简介"
			}
		}
		$(".popup").show();
        var html = template('layui-edit-popup', data);
        document.getElementsByClassName('popup-box')[0].innerHTML = html;
        layui.use('layedit', function(){
		  var layedit = layui.layedit;
		  layedit.build('demo'); //建立编辑器
		});
	});
	
	
	
	
	
	//上传文件
	var isIE = /msie/i.test(navigator.userAgent) && !window.opera;
	function fileChange(target, id) {
		var dataType = $(this).attr("data-file")
		var fileSize = 0;
		var filetypes;
		var filemaxsize;
//		var filemaxsize = 1024 * 2; //2M 
		if(dataType==="video"){
			filetypes = [".rmvb",".rm",".flv",".mp4",".avi"];
			filemaxsize = POSITIVE_INFINITY;
		}
//		var filetypes = [".jpg", ".png", ".rar", ".txt", ".zip", ".doc", ".ppt", ".xls", ".pdf", ".docx", ".xlsx"];
		var filepath = target.value;
		
		if (filepath) {
			var isnext = false;
			var fileend = filepath.substring(filepath.indexOf("."));
			if (filetypes && filetypes.length > 0) {
				for (var i = 0; i < filetypes.length; i++) {
					if (filetypes[i] == fileend) {
						isnext = true;
						break;
					}
				}
			}
			if (!isnext) {
				alert("不接受此文件类型！");
				target.value = "";
				return false;
			}
		} else {
			return false;
		}
		if (isIE && !target.files) {
			var filePath = target.value;
			var fileSystem = new ActiveXObject("Scripting.FileSystemObject");
			if (!fileSystem.FileExists(filePath)) {
				alert("附件不存在，请重新输入！");
				return false;
			}
			var file = fileSystem.GetFile(filePath);
			fileSize = file.Size;
		} else {
			fileSize = target.files[0].size;
		}
		var size = fileSize / 1024;
		if (size > filemaxsize) {
			alert("附件大小不能大于" + filemaxsize / 1024 + "M！");
			target.value = "";
			return false;
		}
		if (size <= 0) {
			alert("附件大小不能为0M！");
			target.value = "";
			return false;
		}
	}
	
	
	/////////发送验证码////////////
	function sendMessage(btn){
		var InterValObj; //timer变量，控制时间
		var count = 60; //间隔函数，1秒执行
		var curCount;//当前剩余秒数
		  　	curCount = count;
		　　//设置button效果，开始计时
		     $(btn).attr("disabled", "true");
		     $(btn).text("请在"+curCount + "秒内输入");
		     InterValObj = window.setInterval(
		     	//timer处理函数
		     	function SetRemainTime() {
		        if (curCount == 0) {                
		            window.clearInterval(InterValObj);//停止计时器
		            $(btn).removeAttr("disabled");//启用按钮
		            $(btn).text("重新发送验证码");
		        }
		        else {
		            curCount--;
		            $(btn).text("请在"+curCount + "秒内输入");
		        }
		    }, 1000); //启动计时器，1秒执行一次
	}

	
	
