<template>
  <div>
    <el-col :span="16" class="contain">
      <el-form :model="form" label-width="160px" ref="form">
        <el-form-item label="礼包名称:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.title}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="礼包描述:">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.content_desc}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="是否随机获取积分">
            <span>
              <el-tag type="gray" v-if="form.content.credits.random_credits">是</el-tag>
              <el-tag type="gray" v-else>否</el-tag>
            </span>
        </el-form-item>
        <el-form-item label="积分随机上限(最多获取)" v-show="form.content.credits.random_credits">
          <el-col :span="24">
            <span>
              <el-tag>{{form.content.credits.upper_limit}}</el-tag>
            </span>
          </el-col>
        </el-form-item>
        <el-form-item label="积分随机下限(最好获取)" v-show="form.content.credits.random_credits">
          <el-col :span="24">
            <span>
              <el-tag>{{form.content.credits.lower_limit}}</el-tag>
            </span>
          </el-col>
        </el-form-item>
        <el-form-item label="赠送固定积分" v-show="!form.content.credits.random_credits">
          <el-col :span="24">
            <el-input v-model.number="form.content.credits.credits"></el-input>
          </el-col>
        </el-form-item>
        <el-form-item label="添加时间:">
          <el-col :span="24">
            <span><el-tag type="gray">{{new Date(form.add_time * 1000).toLocaleString()}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券:">
          <div v-if="form.content.tickets.tickets.length > 0" id="showTickets">
            <el-table :data="form.content.tickets.tickets" style="width: 100%" :row-class-name="tableRowClassName">
              <el-table-column prop="use_base" label="减满基数" width="100"></el-table-column>
              <el-table-column prop="use_sub" label="减满额" width="100"></el-table-column>
              <el-table-column prop="restrict_type_label" label="限制彩种" :formatter="formatter">
              </el-table-column>
              <el-table-column label="开始时间">
                <template scope="scope">
                  <span>{{new Date(scope.row.valid_start * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
              <el-table-column label="结束时间">
                <template scope="scope">
                  <span>{{new Date(scope.row.valid_end * 1000).toLocaleString()}}</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
          <div v-else>
            <span><el-tag type="gray">无</el-tag></span>
          </div>
        </el-form-item>
        <el-form-item label="是否随机获取购彩券">
          <el-col :span="24">
            <!--<el-switch  v-model="form.content.tickets.random_tickets" on-text="是"  off-text="否"></el-switch>-->
            <span>
              <el-tag type="gray" v-if="form.content.tickets.random_tickets">是</el-tag>
               <el-tag type="gray" v-else>否</el-tag>
            </span>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券随机生成上限" v-show="form.content.tickets.random_tickets">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.content.tickets.upper_limit}}</el-tag></span>
          </el-col>
        </el-form-item>
        <el-form-item label="购彩券随机生成下限" v-show="form.content.tickets.random_tickets">
          <el-col :span="24">
            <span><el-tag type="gray">{{form.content.tickets.lower_limit}}</el-tag></span>
          </el-col>
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

  export default {
    data() {
      return {
        lottery_type_list: [],
        gift_package_options: [],      //礼包类型选项.
        form: {
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

        limit_lottery: [],             //限制彩种.

        lotteryList: [],               //彩种列表.

      };
    },
    methods: {

      //获取礼包的类型列表.
      GetGiftTypeList: function () {
        giftApi.giftTypeList().then((res) => {
          this.gift_package_options = res.data.msg;
        })
      },

      //获取彩种类型.
      GetLotteryList: function () {
        lotteryApi.lotteryList().then((res) => {
          let result = res.data.msg;

          for (let i = 0; i < result.length; i++) {
            let obj = {
              value: result[i].Id,
              label: result[i].Name,
            };
            this.lotteryList.push(obj)
          }

          this.options = [{
            value: 1,
            label: "通用",
            children: [{
              value: 0,
              label: "全部彩种",
            }]
          }, {
            value: 2,
            label: "可用彩种",
            children: this.lotteryList
          }, {
            value: 3,
            label: "不可用彩种",
            children: this.lotteryList
          }];

        })
      },

      //获取礼包详细.
      GetGiftDetail: function () {
        let id = this.$route.params.id;
        giftApi.getGiftDetail({id: id}).then((res) => {
          let obj = res.data.msg;

          let tickets = obj.content.tickets.tickets;

          console.log("tickets:", tickets);
//          obj.content.tickets.tickets = tickets;

          let options = this.gift_package_options;
          for (let i = 0; i < options.length; i++) {
            if (options[i].type == obj.gift_type) {
              obj.gift_type = options[i].type_name;
            }
          }

          this.form = obj;
        })
      },

      //根据彩票id获取彩票名称.
      getLotteryNameById: function (id) {
        for (let i = 0; i < this.lotteryList.length; i++) {
          if (this.lotteryList[i].value == id) {
            return this.lotteryList[i].label;
          }
        }
      },
      tableRowClassName(row, index) {
        if (index === 1) {
          return 'info-row';
        } else if (index === 3) {
          return 'positive-row';
        }
        return '';
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
      formatter: function (row, column) {

        let restrict_type = row.restrict_type;
        let restrict_id = row.restrict_id;

        if (restrict_type == 0 && restrict_id == 0) {
          return "通用"
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
      }
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

  .el-table .info-row {
    background: #c9e5f5;
  }

  .el-table .positive-row {
    background: #e2f0e4;
  }
</style>

