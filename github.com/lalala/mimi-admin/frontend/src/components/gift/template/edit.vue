<template>
  <div>
    <el-col :span="16" class="contain">
      <el-form :model="form" label-width="160px" :rules="formRules" ref="form">
        <el-form-item label="礼包名称" prop="title">
          <el-col :span="24">
            <el-input v-model="form.title" auto-complete="off"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="礼包描述" prop="content_desc">
          <el-col :span="24">
            <el-input v-model="form.content_desc"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="是否随机获取积分">
          <el-switch @change="randomCredits" v-model="form.content.credits.random_credits" on-text="是"  off-text="否">
          </el-switch>
        </el-form-item>
        <el-form-item label="积分随机上限(最多获取)" v-show="form.content.credits.random_credits">
          <el-col :span="24">
            <el-input v-model.number="form.content.credits.upper_limit"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="积分随机下限(最好获取)" v-show="form.content.credits.random_credits">
          <el-col :span="24">
            <el-input v-model.number="form.content.credits.lower_limit"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="赠送固定积分" v-show="!form.content.credits.random_credits">
          <el-col :span="24">
            <el-input v-model.number="form.content.credits.credits"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券">
          <div id="add">
            <div id="add_ticket_black" class="form_item" style="display: none">
              <el-row>
                <div class="form_item">
                  <el-form-item label="减满基数">
                    <el-col :span="18">
                      <el-input type="number" v-model.number="ticket.use_base" name="use_base"
                                placeholder="减满基数"></el-input>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="减满额">
                    <el-col :span="18">
                      <el-input type="number" v-model.number="ticket.use_sub" name="use_sub"
                                placeholder="减满额"></el-input>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="限制彩种">
                    <el-col :span="18">
                      <el-cascader :options="options" style="width: 100%" v-model="limit_lottery"
                                   @change="handleChange"></el-cascader>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="开始时间">
                    <el-col :span="18">
                      <el-date-picker v-model="ticket.valid_start" style="width: 100%" type="datetime"
                                      placeholder="选择开始日期时间"></el-date-picker>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="结束时间">
                    <el-col :span="18">
                      <el-date-picker v-model="ticket.valid_end" style="width: 100%" type="datetime"
                                      placeholder="选择结束日期时间"></el-date-picker>
                    </el-col>
                  </el-form-item>
                </div>
              </el-row>
            </div>

            <el-button type="primary" @click="addTicket" id="addBtn">添加购彩券</el-button>
            <el-button id="canclAddBtn" style="display: none" type="primary" @click="canclAddTicket">取消添加</el-button>
            <el-button id="addTicketOK" style="display: none" type="warning" @click="appendTickets">确定添加</el-button>
          </div>

          <div id="edit">
            <div class="form_item" id="edit_ticket_black" style="display: none">
              <el-row>
                <div class="form_item">
                  <el-form-item label="减满基数">
                    <el-col :span="18">
                      <el-input type="number" v-model.number="editTicket.use_base" name="use_base"
                                placeholder="减满基数"></el-input>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="减满额">
                    <el-col :span="18">
                      <el-input type="number" v-model.number="editTicket.use_sub" name="use_sub"
                                placeholder="减满额"></el-input>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="限制彩种">
                    <el-col :span="18">
                      <el-cascader :options="options" style="width: 100%" v-model="limit_lottery"
                                   @change="handleChange"></el-cascader>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="开始时间">
                    <el-col :span="18">
                      <el-date-picker v-model="editTicket.valid_start" style="width: 100%" type="datetime"
                                      placeholder="选择开始日期时间"></el-date-picker>
                    </el-col>
                  </el-form-item>
                </div>
                <div class="form_item">
                  <el-form-item label="结束时间">
                    <el-col :span="18">
                      <el-date-picker v-model="editTicket.valid_end" style="width: 100%" type="datetime"
                                      placeholder="选择结束日期时间"></el-date-picker>
                    </el-col>
                  </el-form-item>
                </div>
              </el-row>
            </div>
            <el-button id="canclEditTicket" style="display: none" type="primary" @click="canclEditTicket">取消修改
            </el-button>
            <el-button id="editTicket" style="display: none" type="warning" @click="editTicketOK">确定修改</el-button>
          </div>
          <div v-if="form.content.tickets.tickets.length > 0" id="showTickets">
            <el-table :data="form.content.tickets.tickets" style="width: 100%">
              <el-table-column prop="use_base" label="减满基数" width="100"></el-table-column>
              <el-table-column prop="use_sub" label="减满额" width="100"></el-table-column>
              <el-table-column prop="restrict_type_label" label="限制彩种" :formatter="formatter"></el-table-column>
              <el-table-column prop="valid_start" label="开始时间" :formatter="formatStartTime"></el-table-column>
              <el-table-column prop="valid_end" label="结束时间" :formatter="formatEndTime"></el-table-column>
              <el-table-column fixed="right" label="操作" width="120">
                <template scope="scope">
                  <el-button type="warning" size="mini" @click="edit(scope.$index)">编辑</el-button>
                  <el-button type="danger" size="mini" @click="del(scope.$index)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-form-item>
        <el-form-item label="是否随机获取购彩券">
          <el-col :span="24">
            <el-switch @change="randomTickets" v-model="form.content.tickets.random_tickets" on-text="是"  off-text="否"></el-switch>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券随机生成上限" v-show="form.content.tickets.random_tickets">
          <el-col :span="24">
            <el-input type="text" v-model.number="form.content.tickets.upper_limit"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券随机生成下限" v-show="form.content.tickets.random_tickets">
          <el-col :span="24">
            <el-input type="text" v-model.number="form.content.tickets.lower_limit"></el-input>
          </el-col>
        </el-form-item>

        <!--<el-form-item label="礼包类型" prop="gift_type" required>-->
          <!--<el-col :span="24">-->
            <!--<el-select v-model.number="form.gift_type" placeholder="请选择" style="width: 100%">-->
              <!--<el-option v-for="item in gift_package_options" :key="item.type" :label="item.type_name"-->
                         <!--:value="item.type"></el-option>-->
            <!--</el-select>-->
          <!--</el-col>-->
        <!--</el-form-item>-->
        <el-form-item>
          <el-button type="primary" @click="onSubmit">确定更改</el-button>
          <router-link :to="{path:'/gift/template/list'}"><el-button>返回</el-button></router-link>
        </el-form-item>
      </el-form>
    </el-col>
  </div>
</template>


<script>

  import giftApi from "../../../api/gift"
  import lotteryApi from "../../../api/lottery"
  import moment from "moment";
  import Vue from 'vue'
  import router from "../../../router"

  export default {
    data() {
      return {
        gift_package_options: [],      //礼包类型选项.
        form: {
          id: "",
          title: "",                   //礼包标题.
          content_desc: "",            //礼包描述.
          gift_type: 0,                //选中的礼包类型.
          content: {
            credits: {
              upper_limit: 0,           //积分随机的上限.
              lower_limit: 0,           //积分随机的下限.
              credits:0,                //积分数.
              random_credits:false,     //是否随机增设积分.
            },                          //赠送的积分.
            tickets: {
              tickets:[],             //具体礼包.
              random_tickets:false,   //是否随机礼包.
              upper_limit: 0,         //礼包随机的上限.
              lower_limit: 0,         //礼包随机的下限.
            },
          },
        },
        ticket: {
          valid_start: "",              //开始时间.
          valid_end: "",                //结束时间.
          use_base: "",                //减满基数.
          use_sub: "",                 //减满额.
          restrict_id: "",              //限制彩种ID.
          restrict_type: ""               //限制的彩票类型，通用=>1, 可用彩种=>2, 不可用彩种=>3
        },

        editTicket: {
          valid_start: "",              //开始时间.
          valid_end: "",                //结束时间.
          use_base: "",                //减满基数.
          use_sub: "",                 //减满额.
          restrict_id: "",              //限制彩种ID.
          restrict_type: ""               //限制的彩票类型，通用=>1, 可用彩种=>2, 不可用彩种=>3
        },
        limit_lottery: [],             //限制彩种.
        lottery_type_list: [],           //所有的彩种类型.
        restrict_type: "",
        restrict_id: "",
        restrict_type_label: "",
        lotteryList: [],               //彩种列表.
        formRules: {
          title: [{required: true, message: '标题不能为空', trigger: 'blur'}],
          content_desc: [{required: true, message: '描述不能为空', trigger: 'blur'}],
          valid_start: [{type: 'date', required: true, message: '请选择开始时间', trigger: 'change'}],
          valid_end: [{type: 'date', required: true, message: '请选择结束时间', trigger: 'change'}],
        },
        options: [],
        status: "add",                   //当前的状态是编辑还是添加.
        tmpTicket: [],                   //临时存储ticket.
        tmpTicketIndex: "",              //临时存储ticket.
      };
    },
    methods: {
      onSubmit() {
        this.disabled = true;
        this.$refs.form.validate((valid) => {
          if (valid) {
            this.$confirm('确认提交吗?', '提示', {type: 'warning'}).then(() => {
                this.form.id = parseInt(this.$route.params.id);
                console.log("this.form:", this.form);
              giftApi.updateGift(this.form).then((res) => {
                this.$message.success("更新成功");
                router.push({path:"/gift/template/list"})
              }).catch(err => {
                this.$message.error("更新失败")
              });
              console.log(this.form)
            }).catch(() => {
              this.disabled = false;
            });
          }
        });
      },

      //赠送积分.
      onCredits: function () {
        let credits = this.form.content.credits;
        if (credits < 1) {
          this.$message.error("赠送的积分不能小于0")
        } else {
          this.dialogVisible = false;
        }
      },


      //获取礼包的类型列表.
      GetGiftTypeList: function () {
        giftApi.giftTypeList().then((res) => {
          this.gift_package_options = res.data.msg;
        })
      },
      addTicket: function () {

        this.ticket.use_base = "";
        this.ticket.use_sub = "";
        this.ticket.valid_start = "";
        this.ticket.valid_end = "";
        this.ticket.restrict_type = "";
        this.limit_lottery = [];
        document.getElementById("add_ticket_black").style = "display:black";
        document.getElementById("canclAddBtn").style = "display:black";
        document.getElementById("addTicketOK").style = "display:black";
        document.getElementById("addBtn").style = "display:none";
      },

      //取消添加.
      canclAddTicket: function () {

        document.getElementById("add_ticket_black").style = "display:none";
        document.getElementById("canclAddBtn").style = "display:none";
        document.getElementById("addTicketOK").style = "display:none";
        document.getElementById("addBtn").style = "display:black";

        //显示数据列表.
        document.getElementById("showTickets").style = "display:black";
      },

      //添加.
      appendTickets: function () {

        let use_base = this.ticket.use_base;
        let use_sub = this.ticket.use_sub;
        let restrict_type = this.restrict_type;
        let valid_start = this.ticket.valid_start;
        let valid_end = this.ticket.valid_end;
        let restrict_type_label = this.restrict_type_label;
        let restrict_id = this.restrict_id;
        //验证减慢基数.
        if (use_base == "") {
          this.$message.error("减满基数不能为空");
          return false
        }

        //验证减慢额.
        if (use_sub == "") {
          this.$message.error("减满额不能为空");
          return false
        }

        let num = use_base / use_sub
        if (num <= 2) {
          this.$message.error("减满额不能超过减满基数的一半");
          return false
        }

        //验证限制彩种.
        if (restrict_type == "" && restrict_type != 0) {
          this.$message.error("限制彩种不能为空");
          return false
        }

        //验证时间.
        if (valid_start == "") {
          this.$message.error("开始时间不能为空");
          return false
        }
        if (valid_end == "") {
          this.$message.error("结束时间不能为空");
          return false
        }
        if (valid_start > valid_end) {
          this.$message.error("开始时间不能大于结束时间");
          return false
        }

        valid_start = moment(valid_start).unix();
        valid_end = moment(valid_end).unix();

        let ticket = {
          use_base: use_base,
          use_sub: use_sub,
          valid_start: valid_start,
          valid_end: valid_end,
          restrict_type_label: restrict_type_label,
          restrict_type: restrict_type,
          restrict_id: restrict_id
        };

        this.form.content.tickets.tickets.push(ticket);

        document.getElementById("add_ticket_black").style = "display:none";
        document.getElementById("canclAddBtn").style = "display:none";
        document.getElementById("addTicketOK").style = "display:none";
        document.getElementById("addBtn").style = "display:black";
      },
      handleChange(value) {
        this.restrict_type = value[0];
        this.restrict_id = value[1];


        let restrict_type = value[0];
        let restrict_id = value[1];

        if (restrict_type == 1) {
          this.restrict_type_label = "全场通用";

          //赋值彩种.
          this.restrict_type = 0;
          this.restrict_id = 0;


        } else if (restrict_type == 2) {

          for (let i = 0; i < this.lotteryList.length; i++) {
            if (this.lotteryList[i].value == restrict_id) {
              this.restrict_type_label = this.lotteryList[i].label + "可用";

              //获取真实的彩票类型.
              this.restrict_type = this.lotteryList[i].type;
              this.restrict_id = this.lotteryList[i].value;

            }
          }

        } else if (restrict_type == 3) {

          for (let i = 0; i < this.lotteryList.length; i++) {
            if (this.lotteryList[i].value == restrict_id) {
              this.restrict_type_label = this.lotteryList[i].label + "不可用";

              //获取真实的彩票类型.
              this.restrict_type = -this.lotteryList[i].type;
              this.restrict_id = -this.lotteryList[i].value;
            }
          }

        } else if (restrict_type == 4) {  //可用类型.

          for (let i = 0; i < this.lottery_type_list.length; i++) {
            if (restrict_id == this.lottery_type_list[i].value) {
              this.restrict_type_label = this.lottery_type_list[i].label + "可用";

              //全部彩种类型可用.
              this.restrict_type = this.lottery_type_list[i].value;
              this.restrict_id = 0
            }
          }

        } else if (restrict_type == 5) {  //不可用类型.
          for (let i = 0; i < this.lottery_type_list.length; i++) {
            if (restrict_id == this.lottery_type_list[i].value) {
              this.restrict_type_label = this.lottery_type_list[i].label + "不可用";

              //全部不可用.
              this.restrict_type = -this.lottery_type_list[i].value;
              this.restrict_id = 0
            }
          }
        }
      },

      //获取彩种类型.
      GetLotteryList: function () {
        lotteryApi.lotteryList().then((res) => {
          let result = res.data.msg;

          for (let i = 0; i < result.length; i++) {
            let obj = {
              value: result[i].Id,
              label: result[i].Name,
              type: result[i].Type,
            };

            this.lotteryList.push(obj)
          }

          this.options = [{
            value: 1,
            label: "全场通用"
          }, {
            value: 2,
            label: "可用彩种",
            children: this.lotteryList
          }, {
            value: 3,
            label: "不可用彩种",
            children: this.lotteryList
          }, {
            value: 4,
            label: "可用类型",
            children: this.lottery_type_list
          }, {
            value: 5,
            label: "不可用类型",
            children: this.lottery_type_list,
          }]
        })
      },
      del: function (index) {
        this.form.content.tickets.tickets.splice(index, 1);
      },

      //编辑.
      edit: function (index) {

        document.getElementById("edit_ticket_black").style = "display:black";
        document.getElementById("showTickets").style = "display:none";
        document.getElementById("canclEditTicket").style = "display:black";
        document.getElementById("editTicket").style = "display:black";
        document.getElementById("addBtn").style = "display:none";


        this.editTicket = this.form.content.tickets.tickets[index];

        let use_base = this.editTicket.use_base;
        let use_sub = this.editTicket.use_sub;
        let restrict_type = this.editTicket.restrict_type;
        let valid_start = this.editTicket.valid_start;
        let valid_end = this.editTicket.valid_end;
        let restrict_type_label = this.restrict_type_label;
        let restrict_id = this.editTicket.restrict_id;
        this.restrict_id = restrict_id;
        this.restrict_type = restrict_type;

        let editTicket = {
          use_base: use_base,
          use_sub: use_sub,
          valid_start: valid_start,
          valid_end: valid_end,
          restrict_type_label: restrict_type_label,
          restrict_type: restrict_type,
          restrict_id: restrict_id
        };

        //彩种限制，回显.
        //如果restrict_type 为 0 并且 restrict_id == 0 则直接选中全场通用.
        //如果restrict_type > 0 并且 restrict_id == 0 则选中可用类型.
        //如果restrict_type < 0 并且 restrict_id == 0 则选中不可用类型.

        //如果restrict_type < 0 并且 restrict_id < 0 则选中不可用彩种.
        //如果restrict_type > 0 并且 restrict_id > 0 则选中可用彩种.

        if (restrict_type == 0 && restrict_id == 0) {
          this.limit_lottery = [1, 0]
        }

        if (restrict_type > 0 && restrict_id == 0) {
          this.limit_lottery = [4, restrict_type]
        }

        if (restrict_type < 0 && restrict_id == 0) {
          this.limit_lottery = [5, -restrict_type]
        }

        if (restrict_type < 0 && restrict_id < 0) {
          this.limit_lottery = [3, -restrict_id]
        }

        if (restrict_type > 0 && restrict_id > 0) {
          this.limit_lottery = [2, restrict_id]
        }

        this.tmpTicket = editTicket;
        this.tmpTicketIndex = index;

        this.editTicket.valid_start = moment(this.editTicket.valid_start * 1000).toDate();
        this.editTicket.valid_end = moment(this.editTicket.valid_end * 1000).toDate();

      },

      //取消修改.
      canclEditTicket: function () {

        this.form.content.tickets.tickets[this.tmpTicketIndex] = this.tmpTicket;

        this.editTicket.valid_start = moment(this.editTicket.valid_start).unix();
        this.editTicket.valid_end = moment(this.editTicket.valid_end).unix();

        document.getElementById("edit_ticket_black").style = "display:none";
        document.getElementById("showTickets").style = "display:black";
        document.getElementById("canclEditTicket").style = "display:none";
        document.getElementById("editTicket").style = "display:none";
        document.getElementById("addBtn").style = "display:black";

        //还原限制彩种.
      },
      editTicketOK: function () {

        let use_base = this.editTicket.use_base;
        let use_sub = this.editTicket.use_sub;
        let restrict_type = this.restrict_type;
        let valid_start = this.editTicket.valid_start;
        let valid_end = this.editTicket.valid_end;
        let restrict_type_label = this.restrict_type_label;
        let restrict_id = this.restrict_id;

        //验证减慢基数.
        if (use_base == "") {
          this.$message.error("减满基数不能为空");
          return false
        }

        //验证减慢额.
        if (use_sub == "") {
          this.$message.error("减满额不能为空");
          return false
        }

        let num = use_base / use_sub;
        if (num <= 2) {
          this.$message.error("减满额不能超过减满基数的一半");
          return false
        }

        //验证限制彩种.
        if (restrict_type == "" && restrict_id != 0) {
          this.$message.error("限制彩种不能为空");
          return false
        }

        //验证时间.
        if (valid_start == "") {
          this.$message.error("开始时间不能为空");
          return false
        }
        if (valid_end == "") {
          this.$message.error("结束时间不能为空");
          return false
        }

        if (valid_start > valid_end) {
          this.$message.error("开始时间不能大于结束时间");
          return false
        }

        valid_start = moment(valid_start).unix();
        valid_end = moment(valid_end).unix();

        let editTicket = {
          use_base: use_base,
          use_sub: use_sub,
          valid_start: valid_start,
          valid_end: valid_end,
          restrict_type_label: restrict_type_label,
          restrict_type: restrict_type,
          restrict_id: restrict_id
        };

//        console.log(this.tmpTicketIndex)

        console.log("editTicket:", editTicket);
        Vue.set(this.form.content.tickets.tickets, this.tmpTicketIndex, editTicket);

        document.getElementById("edit_ticket_black").style = "display:none";
        document.getElementById("showTickets").style = "display:black";
        document.getElementById("canclEditTicket").style = "display:none";
        document.getElementById("editTicket").style = "display:none";
        document.getElementById("addBtn").style = "display:black";

        console.log("this.form.tickets:",this.form.content.tickets)

      },
      formatEndTime: function (row, column) {
        if (typeof(row.valid_end) == "number") {
          return moment.unix(row.valid_end).format('YYYY-MM-DD HH:mm:ss');
        } else {
          let unixTime = moment(row.valid_end).unix();
          return moment.unix(row.unixTime).format('YYYY-MM-DD HH:mm:ss');
        }
      },
      formatStartTime: function (row, column) {
        if (typeof(row.valid_start) == "number") {
          return moment.unix(row.valid_start).format('YYYY-MM-DD HH:mm:ss');
        } else {
          let unixTime = moment(row.valid_start).unix();
          return moment.unix(row.unixTime).format('YYYY-MM-DD HH:mm:ss');
        }
      },

      //获取彩票类型 高频彩，低频彩，竞彩.
      GetLotteryType: function () {
        giftApi.getLotteryTypeList().then((res) => {

          for (let i = 0; i < res.data.msg.length; i++) {
            let obj = {};
            obj.value = res.data.msg[i].type;
            obj.label = res.data.msg[i].name;
            this.lottery_type_list.push(obj)
          }
        })
      },
      //获取礼包详细.
      GetGiftDetail: function () {
        let id = this.$route.params.id;
        giftApi.getGiftDetail({id: id}).then((res) => {
          let obj = res.data.msg;

          let tickets = obj.content.tickets.tickets;

//          obj.tickets = tickets;

          let options = this.gift_package_options;
          for (let i = 0; i < options.length; i++) {
            if (options[i].type == obj.gift_type) {
              obj.gift_type = options[i].type;
            }
          }

          this.form = obj;
        })
      },
      formatter: function (row, column) {

        let restrict_type = row.restrict_type;
        let restrict_id = row.restrict_id;

        if (restrict_type == 0 && restrict_id == 0) {
          return "全场通用"
        }

        if (restrict_type > 0 && restrict_id == 0) {
          //获取类型.
          for (let i = 0; i < this.lottery_type_list.length; i++) {
            if (this.lottery_type_list[i].value == restrict_type) {
              return this.lottery_type_list[i].label + "可用";
            }
          }
        }

        if (restrict_type < 0 && restrict_id == 0) {
          for (let i = 0; i < this.lottery_type_list.length; i++) {
            if (this.lottery_type_list[i].value == -restrict_type) {
              return this.lottery_type_list[i].label + "不可用";
            }
          }
        }

        if (restrict_type < 0 && restrict_id < 0) {
          //获取彩种.
          return this.getLotteryNameById(-restrict_id) + "不可用";
        }

        if (restrict_type > 0 && restrict_id > 0) {
          //获取彩种.
          return this.getLotteryNameById(restrict_id) + "可用";
        }
      },
      //根据彩票id获取彩票名称.
      getLotteryNameById: function (id) {
        for (let i = 0; i < this.lotteryList.length; i++) {
          if (this.lotteryList[i].value == id) {
            return this.lotteryList[i].label;
          }
        }
      },
      //是否随机增设积分.
      randomCredits:function (val) {
        if (this.val) {

        }
      },

      //是否随机生成购彩券.
      randomTickets:function (val) {

      },
    },
    mounted() {
      this.GetLotteryType();
      this.GetGiftDetail();
    }
  }
</script>

<style>
  .contain {
    margin: 50px;
  }

  .form_item {
    margin-top: 20px;
    margin-bottom: 20px;
    background: #f6f6f6;
    border-radius: 3px;
  }

  .ri {
    text-align: right;
  }

  .specifyTypeShow {
    display: none;
  }

  .specifyTypeNotShow {
    display: block;
  }

  #showTickets, #edit_ticket_black {
    margin-top: 15px;
    margin-bottom: 15px;
  }
</style>

