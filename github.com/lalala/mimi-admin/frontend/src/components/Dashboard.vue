<template>
    <div class="cont">
        <el-row :gutter="24">
          <el-col :span="12">
              <el-col><h3 style="margin-left:0;padding-left:0;">收入</h3></el-col>
              <el-col :span="12">
                <div class="income box">
                    <el-col :span="6">
                        <div class="pic bg1">¥</div>
                    </el-col>
                    <el-col :span="18">
                        <div class="income_header" style="margin-top:20px;"><h5>¥ 6542</h5></div>
                        <div class="income_header"><h6>今日收入</h6></div>
                    </el-col>
                </div>
              </el-col>
              <el-col :span="12">
                <div class="income box">
                    <el-col :span="6">
                        <div class="pic bg2">¥</div>
                    </el-col>
                    <el-col :span="18">
                        <div class="income_header" style="margin-top:20px;"><h5>¥ 956.132.10</h5></div>
                        <div class="income_header"><h6>历史总收入</h6></div>
                    </el-col>
                </div>
              </el-col>
          </el-col>
          <el-col :span="12">
              <el-col><h3>订单</h3></el-col>
              <el-col :span="12">
                  <div class="income box">
                     <el-col :span="6">
                        <div class="pic bg3">订</div>
                     </el-col>
                     <el-col :span="18">
                        <div class="income_header" style="margin-top:20px;"><h5>7562</h5></div>
                        <div class="income_header"><h6>新增订单</h6></div>
                     </el-col>
                     </div>
              </el-col>
              <el-col :span="12">
                    <div class="income box">
                        <el-col :span="6">
                            <div class="pic bg4">订</i></div>
                        </el-col>
                        <el-col :span="18">
                            <div class="income_header" style="margin-top:20px;"><h5>767547</h5></div>
                            <div class="income_header"><h6>历史总订单</h6></div>
                        </el-col>
                    </div>
              </el-col>
          </el-col>
          <el-col :span="12" >
            <el-col><h3  style="margin-left:0;padding-left:0;">访问曲线图</h3></el-col>
             <el-col :span="24" style="">
                <div id="main" style="width: 100%;height:400px;" class="box"></div>
             </el-col>
          </el-col>
          <el-col :span="12">
            <el-col><h3>订单曲线图</h3></el-col>

                     <el-col :span="24">
                        <div id="order" class="box" style="height:400px;"></div>
                     </el-col>

          </el-col>
          <el-col :span="12">
            <el-col><h3>访问地理分布图</h3></el-col>
            <el-col :span="24"><div id="map" class="box" style="height:890px;"></div></el-col>
          </el-col>
          <el-col :span="12">
              <el-col><h3>用户</h3></el-col>
                  <el-col :span="12">
                    <div class="box">
                      <div class="grid-content bg-purple block_top bg bg2">新增用户</div>
                      <div class="grid-content bg-purple block bg2">2000</div>
                      </div>
                    </el-col>
                    <el-col :span="12">
                        <div class="box">
                          <div class="grid-content bg-purple block_top bg bg3">总用户数</div>
                          <div class="grid-content bg-purple block bg3">36532</div>
                        </div>
                    </el-col>
          </el-col>
          <el-col :span="12">
            <el-col><h3>用户留存率</h3></el-col>
            <el-col :span="24">
            <div class="box">
                <el-table :data="table" border style="width: 100%">
                    <el-table-column prop="date" label="首次使用时间" width="150"></el-table-column>
                        <el-table-column prop="new_user_num" label="新增用户数" width="120"></el-table-column>
                        <el-table-column label="留存率" style="text-align:center;">
                          <el-table-column label="第一天" prop="day1"></el-table-column>
                          <el-table-column prop="day2" label="第二天"></el-table-column>
                          <el-table-column prop="day3" label="第三天"></el-table-column>
                          <el-table-column prop="day4" label="第四天"></el-table-column>
                          <el-table-column prop="day5" label="第五天"></el-table-column>
                          <el-table-column prop="day6" label="第六天"></el-table-column>
                          <el-table-column prop="day7" label="第七天"></el-table-column>
                      </el-table-column>
                    </el-table>
                </div>
                </el-col>
          </el-col>
          <el-col :span="12">
                <el-col><h3>短信</h3></el-col>
                <el-col :span="12">
                    <div id="mobile_message" class="box" style="width: 100%;height:400px;"></div>
                </el-col>
                  <el-col :span="12">
                    <div  id="mobile_message_type" class="box" style="width: 100%;height:400px;"></div>
                  </el-col>
            </el-col>
            <el-col :span="24">
                <el-col :span="12" style="padding-left:0;">
                  <el-col><h3>彩票中奖人数比例图</h3></el-col>
                    <el-col :span="24">
                        <div id="lottery_winning" class="box" style="width:100%;height:400px;"></div>
                    </el-col>
                  </el-col>
                   <el-col :span="12">
                        <el-col><h3>彩票收入比例图</h3></el-col>
                        <el-col :span="24">
                            <div id="lottery" class="box" style="width:100%;height:400px;"></div>
                        </el-col>
                   </el-col>
            </el-col>
            <el-col :span="12">
                <el-col><h3>彩票订单比例图</h3></el-col>
                <el-col :span="24">
                    <div id="lottery_order" class="box" style="width:100%;height:400px;"></div>
                </el-col>
            </el-col>
            <el-col :span="12">
                <el-col><h3>彩票中奖金额比例图</h3></el-col>
                <el-col :span="24">
                    <div id="lottery_amount" class="box" style="width:100%;height:400px;"></div>
                </el-col>
            </el-col>
        </el-row>
    </div>
</template>

<style>
    .cont {
        min-width:1670px;
    }
    body {
        background:#f2f4f8;
    }
    h3{
        margin-top:20px;
        margin-left:0;
    }
    h5 {
        font-size: 2em;
        font-weight: bold;
        padding:0;
        margin:0;
        color:#333333;
    }
    h6 {
        font-size: 1rem;
        font-weight: bold;
        padding:0;
        margin:0;
        color:#cccccc;
    }
    .block_top {
        border-radius:3px 3px 0 0;
        color:#ffffff;
        font-size:1.5em;
        text-align:center;
        padding-top:10px;
        padding-bottom:10px;
    }
    .block {
        border-radius:0 0 3px 3px;
        color:#ffffff;
        font-size:2em;
        text-align:center;
        height:80px;
        line-height:80px;
    }
    .bg {
        box-shadow: inset 0 -50px 0 0 rgba(255,255,255,.1);
        border-top-right-radius: 3px;
        border-top-left-radius: 3px;
        text-transform: uppercase;
    }
    .bg1 {
        background:#ef553a;
    }
    .bg2 {
        background:#26c281;
    }
    .bg3 {
        background:#57889c;
    }
    .bg4 {
        background:#9358ac
    }
    .bg5 {
        background:#3b5998;
    }
    .bg6 {
        background:#00aced;
    }
    .bg7 {
        background:#FF6666;
    }
    .cell {
        text-align:center;
    }
    .income {
        width:100%;
        height:100px;
    }
    .box {
        box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
        -webkit-box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.2);
    }
    .box:hover {
        box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.5);
        -webkit-box-shadow: 0 1px 3px 0px rgba(0, 0, 0, 0.5);
        transition: box-shadow 0.8s;
    }
    .pic {
        width:80px;
        height:80px;
        border-radius:50px;
        margin-top:10px;
        text-align:center;
        font-size:3em;
        line-height:80px;
        color:#ffffff;
    }
    .income_header {
        width:100%;
        height:30px;
        line-height:30px;

    }
</style>

<script>
    export default {
        data() {
            return {
                data : [],
                visit_quantity : [{
                    date : "2017-04-05",
                    uv : "2000",
                    ip : "1200",
                    new_uv : "400",
                },{
                    date : "2017-04-04",
                    uv : "2432",
                    ip : "1000",
                    new_uv : "100",
                },{
                  date : "2017-04-03",
                  uv : "2432",
                  ip : "1000",
                  new_uv : "100",
                },{
                  date : "2017-04-02",
                  uv : "2432",
                  ip : "1000",
                  new_uv : "100",
                },{
                    date : "2017-04-01",
                    uv : "2432",
                    ip : "1000",
                    new_uv : "100",
                },{
                  date : "2017-03-31",
                  uv : "2432",
                  ip : "1000",
                  new_uv : "100",
                }],
                table : [{
                      date: '2017-04-03',
                      new_user_num: 3234,
                      day1: "90%",
                      day2: "80%",
                      day3: "70%",
                      day4: "60%",
                    },{
                        date: '2017-04-04',
                        new_user_num: 4234,
                        day1: "90%",
                        day2: "80%",
                        day3: "70%",
                    },{
                        date: '2017-04-05',
                        new_user_num: 1345,
                        day1: "90%",
                        day2: "80%",
                     },{
                         date: '2017-04-06',
                         new_user_num: 1000,
                         day1: "100%",
                     }]
            }
        },methods : {
            getLotteryListInfo() {

            },
            access () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');

                //基于准备好的dom，初始化echarts实例
                var myChart = echarts.init(document.getElementById('main'));

                function randomData() {
                    now = new Date(+now + oneDay);
                    value = value + Math.random() * 21 - 10;
                    return {
                        name: now.toString(),
                        value: [
                            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
                            Math.round(value)
                        ]
                    }
                }

                var data = [];
                var now = new Date(2017, 4, 6);
                var oneDay = 24 * 3600 * 1000;
                var value = Math.random() * 1000;
                for (var i = 0; i < 1000; i++) {
                    data.push(randomData());
                }
                var option = {
                        tooltip : {
                            trigger: 'axis',
                            axisPointer: {
                                type: 'cross',
                                label: {
                                    backgroundColor: '#6a7985'
                                }
                            }
                        },
                        legend: {
                            data:['Android','iOS']
                        },
                        grid: {
                            left: '3%',
                            right: '4%',
                            bottom: '3%',
                            containLabel: true
                        },
                        xAxis : [
                            {
                                type : 'category',
                                boundaryGap : false,
                                data : ['周一','周二','周三','周四','周五','周六','周日']
                            }
                        ],
                        yAxis : [
                            {
                                type : 'value'
                            }
                        ],
                        series : [
                            {
                                name:'iOS访问',
                                type:'line',
                                stack: '总量',
                                areaStyle: {normal: {}},
                                data:[120, 132, 101, 134, 90, 230, 210]
                            },
                            {
                                name:'Android访问',
                                type:'line',
                                stack: '总量',
                                areaStyle: {normal: {}},
                                data:[220, 182, 191, 234, 290, 330, 310]
                            },
                        ],
                        color : ['#61a0a8', '#d48265', '#91c7ae','#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3'],
                    };


                        // 使用刚指定的配置项和数据显示图表。
                        myChart.setOption(option);
            },
            order () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');

                //基于准备好的dom，初始化echarts实例
                var orderChart = echarts.init(document.getElementById('order'));

                function randomData() {
                    now = new Date(+now + oneDay);
                    value = value + Math.random() * 21 - 10;
                    return {
                        name: now.toString(),
                        value: [
                            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
                            Math.round(value)
                        ]
                    }
                }

                var data = [];
                var now = new Date(2017, 4, 6);
                var oneDay = 24 * 3600 * 1000;
                var value = Math.random() * 1000;
                for (var i = 0; i < 1000; i++) {
                    data.push(randomData());
                }

                var option = {
                        tooltip: {
                            trigger: 'axis',
                            formatter: function (params) {
                                params = params[0];
                                var date = new Date(params.name);
                                return date.getDate() + '/' + (date.getMonth() + 1) + '/' + date.getFullYear() + ' : ' + params.value[1];
                            },
                            axisPointer: {
                                animation: false
                            }
                        },
                        xAxis: {
                            type: 'time',
                            splitLine: {
                                show: false
                            }
                        },
                        yAxis: {
                            type: 'value',
                            boundaryGap: [0, '100%'],
                            splitLine: {
                                show: false
                            }
                        },
                        series: [{
                            name: '模拟数据',
                            type: 'line',
                            showSymbol: false,
                            hoverAnimation: false,
                            data: data
                        }],
                        color : ['#d48265', '#91c7ae','#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3'],
                    };

                        // 使用刚指定的配置项和数据显示图表。
                        orderChart.setOption(option);
            },
            mobile () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');

                //基于准备好的dom，初始化echarts实例
                var mobileChart = echarts.init(document.getElementById('mobile_message'));

                function randomData() {
                    now = new Date(+now + oneDay);
                    value = value + Math.random() * 21 - 10;
                    return {
                        name: now.toString(),
                        value: [
                            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
                            Math.round(value)
                        ]
                    }
                }

                var data = [];
                var now = new Date(2017, 4, 6);
                var oneDay = 24 * 3600 * 1000;
                var value = Math.random() * 1000;
                for (var i = 0; i < 1000; i++) {
                    data.push(randomData());
                }

                var option = {
                     tooltip : {
                         trigger: 'item',
                         formatter: "{a} <br/>{b} : {c} ({d}%)"
                     },
                     legend: {
                         orient: 'vertical',
                         left: 'left',
                         data: ['已发送','剩余条数']
                     },
                     series : [
                         {
                             name: '短信条数',
                             type: 'pie',
                             radius : '55%',
                             center: ['50%', '60%'],
                             data:[
                                 {value:1000, name:'已发送'},
                                 {value:500, name:'剩余条数'},
                             ],
                             itemStyle: {
                                 emphasis: {
                                     shadowBlur: 10,
                                     shadowOffsetX: 0,
                                     shadowColor: 'rgba(0, 0, 0, 0.5)'
                                 }
                             }
                         }
                     ],
                     color : ['#d48265', '#91c7ae','#749f83',  '#ca8622', '#bda29a','#6e7074', '#546570', '#c4ccd3'],
                 };

                // 使用刚指定的配置项和数据显示图表。
                mobileChart.setOption(option);
            },
            mobileType () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');

                //基于准备好的dom，初始化echarts实例
                var mobileTypeChart = echarts.init(document.getElementById('mobile_message_type'));

                function randomData() {
                    now = new Date(+now + oneDay);
                    value = value + Math.random() * 21 - 10;
                    return {
                        name: now.toString(),
                        value: [
                            [now.getFullYear(), now.getMonth() + 1, now.getDate()].join('/'),
                            Math.round(value)
                        ]
                    }
                }

                var data = [];
                var now = new Date(2017, 4, 6);
                var oneDay = 24 * 3600 * 1000;
                var value = Math.random() * 1000;
                for (var i = 0; i < 1000; i++) {
                    data.push(randomData());
                }

                var option = {
                     tooltip : {
                         trigger: 'item',
                         formatter: "{a} <br/>{b} : {c} ({d}%)"
                     },
                     legend: {
                         orient: 'vertical',
                         left: 'left',
                         data: ['注册验证码','找回密码','修改密码','广告通知']
                     },
                     series : [
                         {
                             name: '短信使用情况',
                             type: 'pie',
                             radius : '55%',
                             center: ['50%', '60%'],
                             data:[
                                 {value:1000, name:'注册验证码'},
                                 {value:342, name:'找回密码'},
                                 {value:100, name:'修改密码'},
                                 {value:5000, name:'广告通知'},
                             ],
                             itemStyle: {
                                 emphasis: {
                                     shadowBlur: 10,
                                     shadowOffsetX: 0,
                                     shadowColor: 'rgba(0, 0, 0, 0.5)'
                                 }
                             }
                         }
                     ],
                     color : ['#bda29a','#6e7074', '#546570', '#c4ccd3'],
                 };

                // 使用刚指定的配置项和数据显示图表。
                mobileTypeChart.setOption(option);
            },
            map () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');
                require('echarts/lib/component/title');

                require('echarts/map/js/china')
                //基于准备好的dom，初始化echarts实例
                var mapChart = echarts.init(document.getElementById('map'));

                function randomData() {
                    return Math.round(Math.random()*1000);
                }

                var option = {
                        tooltip : {
                            trigger: 'item'
                        },
                        dataRange: {
                            min: 0,
                            max: 2500,
                            x: 'left',
                            y: 'bottom',
                            text:['高','低'],           // 文本，默认为数值文本
                            calculable : true
                        },
                        toolbox: {
                            show: true,
                            orient : 'vertical',
                            x: 'right',
                            y: 'center',
                            feature : {
                                mark : {show: true},
                                dataView : {show: true, readOnly: false},
                                restore : {show: true},
                                saveAsImage : {show: true}
                            }
                        },
                        roamController: {
                            show: true,
                            x: 'right',
                            mapTypeControl: {
                                'china': true
                            }
                        },
                        series : [
                            {
                                name: '地区访问情况',
                                type: 'map',
                                mapType: 'china',
                                roam: false,
                                itemStyle:{
                                    normal:{label:{show:true}},
                                    emphasis:{label:{show:true}}
                                },
                                data:[
                                    {name: '北京',value: Math.round(Math.random()*1000)},
                                    {name: '天津',value: Math.round(Math.random()*1000)},
                                    {name: '上海',value: Math.round(Math.random()*1000)},
                                    {name: '重庆',value: Math.round(Math.random()*1000)},
                                    {name: '河北',value: Math.round(Math.random()*1000)},
                                    {name: '河南',value: Math.round(Math.random()*1000)},
                                    {name: '云南',value: Math.round(Math.random()*1000)},
                                    {name: '辽宁',value: Math.round(Math.random()*1000)},
                                    {name: '黑龙江',value: Math.round(Math.random()*1000)},
                                    {name: '湖南',value: Math.round(Math.random()*1000)},
                                    {name: '安徽',value: Math.round(Math.random()*1000)},
                                    {name: '山东',value: Math.round(Math.random()*1000)},
                                    {name: '新疆',value: Math.round(Math.random()*1000)},
                                    {name: '江苏',value: Math.round(Math.random()*1000)},
                                    {name: '浙江',value: Math.round(Math.random()*1000)},
                                    {name: '江西',value: Math.round(Math.random()*1000)},
                                    {name: '湖北',value: Math.round(Math.random()*1000)},
                                    {name: '广西',value: Math.round(Math.random()*1000)},
                                    {name: '甘肃',value: Math.round(Math.random()*1000)},
                                    {name: '山西',value: Math.round(Math.random()*1000)},
                                    {name: '内蒙古',value: Math.round(Math.random()*1000)},
                                    {name: '陕西',value: Math.round(Math.random()*1000)},
                                    {name: '吉林',value: Math.round(Math.random()*1000)},
                                    {name: '福建',value: Math.round(Math.random()*1000)},
                                    {name: '贵州',value: Math.round(Math.random()*1000)},
                                    {name: '广东',value: Math.round(Math.random()*1000)},
                                    {name: '青海',value: Math.round(Math.random()*1000)},
                                    {name: '西藏',value: Math.round(Math.random()*1000)},
                                    {name: '四川',value: Math.round(Math.random()*1000)},
                                    {name: '宁夏',value: Math.round(Math.random()*1000)},
                                    {name: '海南',value: Math.round(Math.random()*1000)},
                                    {name: '台湾',value: Math.round(Math.random()*1000)},
                                    {name: '香港',value: Math.round(Math.random()*1000)},
                                    {name: '澳门',value: Math.round(Math.random()*1000)}
                                ]
                            }
                        ]
                    };


                // 使用刚指定的配置项和数据显示图表。
                mapChart.setOption(option);
            },
            lottery () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');
                require('echarts/lib/component/title');

                require('echarts/map/js/china')
                //基于准备好的dom，初始化echarts实例
                var lotteryChart = echarts.init(document.getElementById('lottery'));


                var option = {
                        tooltip: {
                            trigger: 'item',
                            formatter: "{a} <br/>{b}: {c} ({d}%)"
                        },
                        legend: {
                            orient: 'vertical',
                            x: 'left',
                            data:['重庆时时彩','双色球','福彩3D','竞猜足球','北京PK拾']
                        },
                        series: [
                            {
                                name:'彩票中奖比例图',
                                type:'pie',
                                radius: ['50%', '70%'],
                                avoidLabelOverlap: false,
                                label: {
                                    normal: {
                                        show: false,
                                        position: 'center'
                                    },
                                    emphasis: {
                                        show: true,
                                        textStyle: {
                                            fontSize: '30',
                                            fontWeight: 'bold'
                                        }
                                    }
                                },
                                labelLine: {
                                    normal: {
                                        show: false
                                    }
                                },
                                data:[
                                    {value:335, name:'重庆时时彩'},
                                    {value:310, name:'双色球'},
                                    {value:234, name:'福彩3D'},
                                    {value:135, name:'竞猜足球'},
                                    {value:1548, name:'北京PK拾'}
                                ]
                            }
                        ]
                    };
                // 使用刚指定的配置项和数据显示图表。
                lotteryChart.setOption(option);
            },
            lottery_winning () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');
                require('echarts/lib/component/title');

                require('echarts/map/js/china')
                //基于准备好的dom，初始化echarts实例
                var lottery_winningChart = echarts.init(document.getElementById('lottery_winning'));


                var option = {
                        tooltip: {
                            trigger: 'item',
                            formatter: "{a} <br/>{b}: {c} ({d}%)"
                        },
                        legend: {
                            orient: 'vertical',
                            x: 'left',
                            data:['重庆时时彩','双色球','福彩3D','竞猜足球','北京PK拾']
                        },
                        series: [
                            {
                                name:'彩票收入比例图',
                                type:'pie',
                                radius: ['50%', '70%'],
                                avoidLabelOverlap: false,
                                label: {
                                    normal: {
                                        show: false,
                                        position: 'center'
                                    },
                                    emphasis: {
                                        show: true,
                                        textStyle: {
                                            fontSize: '30',
                                            fontWeight: 'bold'
                                        }
                                    }
                                },
                                labelLine: {
                                    normal: {
                                        show: false
                                    }
                                },
                                data:[
                                    {value:651, name:'重庆时时彩'},
                                    {value:785, name:'双色球'},
                                    {value:365, name:'福彩3D'},
                                    {value:544, name:'竞猜足球'},
                                    {value:899, name:'北京PK拾'}
                                ]
                            }
                        ]
                    };
                // 使用刚指定的配置项和数据显示图表。
                lottery_winningChart.setOption(option);
            },
            lottery_order () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');
                require('echarts/lib/component/title');

                require('echarts/map/js/china')
                //基于准备好的dom，初始化echarts实例
                var lotteryOrderChart = echarts.init(document.getElementById('lottery_order'));


                var option = {
                        tooltip: {
                            trigger: 'item',
                            formatter: "{a} <br/>{b}: {c} ({d}%)"
                        },
                        legend: {
                            orient: 'vertical',
                            x: 'left',
                            data:['重庆时时彩','双色球','福彩3D','竞猜足球','北京PK拾']
                        },
                        series: [
                            {
                                name:'彩票收入比例图',
                                type:'pie',
                                radius: ['50%', '70%'],
                                avoidLabelOverlap: false,
                                label: {
                                    normal: {
                                        show: false,
                                        position: 'center'
                                    },
                                    emphasis: {
                                        show: true,
                                        textStyle: {
                                            fontSize: '30',
                                            fontWeight: 'bold'
                                        }
                                    }
                                },
                                labelLine: {
                                    normal: {
                                        show: false
                                    }
                                },
                                data:[
                                    {value:66, name:'重庆时时彩'},
                                    {value:785, name:'双色球'},
                                    {value:212, name:'福彩3D'},
                                    {value:5465, name:'竞猜足球'},
                                    {value:2132, name:'北京PK拾'}
                                ]
                            }
                        ]
                    };
                // 使用刚指定的配置项和数据显示图表。
                lotteryOrderChart.setOption(option);
            },
            lottery_amount () {

                // 引入 ECharts 主模块
                var echarts = require('echarts/lib/echarts');
                // 引入柱状图
                require('echarts/lib/chart/bar');
                // 引入提示框和标题组件
                require('echarts/lib/component/tooltip');
                require('echarts/lib/component/title');
                require('echarts/lib/component/title');

                require('echarts/map/js/china')
                //基于准备好的dom，初始化echarts实例
                var lotteryAmountChart = echarts.init(document.getElementById('lottery_amount'));


                var option = {
                        tooltip: {
                            trigger: 'item',
                            formatter: "{a} <br/>{b}: {c} ({d}%)"
                        },
                        legend: {
                            orient: 'vertical',
                            x: 'left',
                            data:['重庆时时彩','双色球','福彩3D','竞猜足球','北京PK拾']
                        },
                        series: [
                            {
                                name:'彩票收入比例图',
                                type:'pie',
                                radius: ['50%', '70%'],
                                avoidLabelOverlap: false,
                                label: {
                                    normal: {
                                        show: false,
                                        position: 'center'
                                    },
                                    emphasis: {
                                        show: true,
                                        textStyle: {
                                            fontSize: '30',
                                            fontWeight: 'bold'
                                        }
                                    }
                                },
                                labelLine: {
                                    normal: {
                                        show: false
                                    }
                                },
                                data:[
                                    {value:66, name:'重庆时时彩'},
                                    {value:785, name:'双色球'},
                                    {value:212, name:'福彩3D'},
                                    {value:5465, name:'竞猜足球'},
                                    {value:2132, name:'北京PK拾'}
                                ]
                            }
                        ]
                    };
                // 使用刚指定的配置项和数据显示图表。
                lotteryAmountChart.setOption(option);
            },
            getAddress () {

            }
        },
        mounted() {
            //this.getLotteryListInfo();
            this.getAddress();
            this.access();
            this.order();
            this.mobile();
            this.mobileType();
            this.map();
            this.lottery();
            this.lottery_winning();
            this.lottery_order();
            this.lottery_amount();
        }
    }

</script>