<template>
<div>
<el-col :span="8" class="contain">
<el-form ref="form" :model="form" label-width="80px">
  <el-form-item label="活动标题" >
    <el-col :span="24">
        <el-input v-model="form.title"></el-input>
    </el-col>
  </el-form-item>
  <el-form-item label="活动描述">
    <el-col :span="24">
     <el-input v-model="form.desc"></el-input>
     </el-col>
  </el-form-item>
  <el-form-item label="活动内容">
    <el-col :span="24">
        <el-input type="textarea" v-model="form.content"></el-input>
    </el-col>
  </el-form-item>
    <el-form-item label="赠送类型">
        <el-radio-group v-model="form.gift_type">
            <el-radio label="兑换码"></el-radio>
            <el-radio label="邀请码"></el-radio>
        </el-radio-group>
    </el-form-item>
    <el-form-item label="活动性质" prop="type">
        <el-checkbox-group v-model="form.type">
            <el-checkbox label="赠送积分" @change="credits()" v-model="checked" name="type"></el-checkbox>
            <el-dialog title="" :visible.sync="dialogVisible" size="tiny" >
                <el-form-item label="赠送积分">
                    <el-col :span="12">
                        <el-input type="number" placeholder="输入赠送的积分额度"></el-input>
                    </el-col>
                </el-form-item>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
                </span>
            </el-dialog>

            <el-checkbox label="购彩券" @change="tickets()" name="type"></el-checkbox>
            <el-dialog title="" :visible.sync="dialogTicketVisible" size="tiny" >
                <el-row  :gutter="24">
                   <div class="form_item">
                        <el-form-item>
                            <el-col :span="4" class="ri">
                                <label>减满基数:</label>
                           </el-col>
                            <el-col :span="20">
                                <el-input  type="number" name="use_base" placeholder="减满基数"></el-input>
                            </el-col>
                        </el-form-item>
                    </div>
                    <div class="form_item">
                        <el-form-item>
                            <el-col :span="4" class="ri">
                                <label>减满额:</label>
                            </el-col>
                            <el-col :span="20">
                                <el-input type="number" name="use_sub" placeholder="减满额"></el-input>
                            </el-col>
                        </el-form-item>
                    </div>
                    <div class="form_item">
                        <el-form-item>
                            <el-col :span="4" class="ri">
                                <label>限制彩种:</label>
                            </el-col>
                            <el-col :span="20">
                                <el-radio-group v-model="form.gift_type2">

                                    <el-form-item>
                                        <el-col><el-radio label="通用" @click.native="blur"></el-radio></el-col>
                                    </el-form-item>

                                    <el-form-item>
                                        <el-col>
                                            <el-radio label="指定彩种类型"  @click.native="specifyType"></el-radio>
                                        </el-col>
                                        <el-col>
                                            <div class="specifyType" v-if="this.IsA == true">
                                                <el-checkbox-group v-model="checkList">
                                                    <el-checkbox v-for="cp in cpList" :label="cp">{{cp}}</el-checkbox>
                                                </el-checkbox-group>
                                            </div>
                                        </el-col>
                                    </el-form-item>

                                    <el-form-item>
                                        <el-col><el-radio label="不可用于指定彩种类型" @click.native="specifyNotType"></el-radio></el-col>
                                        <el-col>
                                            <div class="specifyNotType" v-if="this.IsB == true">
                                                <el-checkbox-group v-model="checkList">
                                                    <el-checkbox v-for="cp in cpList" :label="cp" >{{cp}}</el-checkbox>
                                                </el-checkbox-group>
                                            </div>
                                        </el-col>
                                    </el-form-item>
                                </el-radio-group>
                            </el-col>
                        </el-form-item>
                    </div>
                </el-row>
            <span slot="footer" class="dialog-footer">
                <el-button @click="ticketClose()">取 消</el-button>
                <el-button type="primary" @click="dialogTicketVisible = false">确 定</el-button>
            </span>
        </el-dialog>

        </el-checkbox-group>
    </el-form-item>
  <el-form-item label="结束时间">
      <el-col :span="24">
        <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
      </el-col>
    </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="onSubmit">立即创建</el-button>
    <el-button>取消</el-button>
  </el-form-item>
</el-form>
</el-col>
</div>
</template>


<script>
    const list = ['双色球','北京PK拾','福彩3D','竞彩足球'];
    export default {
      data() {
        return {
           IsA:false,
           IsB:false,
           dialogVisible: false,
           dialogTicketVisible:false,
           form: {
             title: '',
             desc: '',
             content: '',
             date: '',
             type: [],
             gift_type : "",
               gift_type2 : [],
           },
            checked : false,
            ticket : false,
            cp_type : 0,
            checkList:[],
            cpList: list,

        }
      },
      methods: {
        onSubmit() {
          console.log('submit!');

        },
          credits() {
            var status = this.checked;
            if (!status) {
                this.checked = true;
                this.dialogVisible = true
            } else {
                this.checked = false;
                this.dialogVisible = false
            }
        },
        tickets() {
            var status = this.ticket;
            if (!status) {
                this.ticket = true;
                this.dialogTicketVisible = true
            } else {
                this.ticket = false;
                this.dialogTicketVisible = false
            }
        },
        ticketClose() {
            this.dialogTicketVisible = false
        },
        specifyType() {
            this.IsA = true;
            this.IsB = false;
        },
        specifyNotType() {
            this.IsA = false;
            this.IsB = true;
        },
          blur() {
            this.IsA = false;
              this.IsB = false;
          }
      }
    }
</script>

<style>
    .contain {
        margin:50px;
    }
    .form_item {
        margin-top:15px;
        margin-bottom:15px;
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
</style>

