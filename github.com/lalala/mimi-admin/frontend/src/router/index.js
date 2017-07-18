import Vue from 'vue'
import Router from 'vue-router'

import NotFound from '@/components/NotFound'
import Login from '@/components/Login'
import Home from '@/components/Home'
import Profile from '@/components/setting/Profile'
import RoleList from '@/components/setting/RoleList'
import UserList from '@/components/setting/UserList'
import PrivList from '@/components/setting/Privilege'
import Demo from '@/components/setting/demo'
import Order from '@/components/order/Order'
import Dashboard from '@/components/Dashboard'
import News from '@/components/discover/NewsList'
import NewsForm from '@/components/discover/NewsForm'
import DiscoverBanner from '@/components/discover/BannerList'

Vue.use(Router);

export default new Router({
  routes: [
    // 404
    {
      path: '/404',
      name: '',
      component: NotFound,
      hidden: true
    },
    // 登录
    {
      path: '/login',
      name: '',
      component: Login,
      hidden: true
    },
    //数据.
    {
      path : "/",
      component:Home,
      name: "数据统计",
      iconCls: 'el-icon-menu',
      children: [
        { path: '/data/dashboard', component: require('@/components/data/dashboard'), name: '控制台' },
        { path: '/data/recharge', component: require('@/components/data/recharge/recharge'), name: '充值统计' },
        { path: '/data/buycai', component: require('@/components/data/buycai/buycai'), name: '购彩统计' },
      ]
    },
    // 系统管理
    {
      path: '/',
      component: Home,
      name: '系统管理',
      iconCls: 'el-icon-setting',
      children: [
        {path: '/dashboard',name: '控制台',component: Dashboard},
        // { path: '/setting/profile', component: Profile, name: '个人中心' },
        // { path: '/setting/userlist', component: UserList, name: '用户列表' },
        { path: '/user/role/list',      component: require('@/components/user/role/list'),      name: '角色列表' },
        { path: '/user/role/add',       component: require('@/components/user/role/add'),       name: '添加角色', hidden:true },
        { path: '/user/privilege/list', component: require('@/components/user/privilege/list'), name: '权限管理', },
        // { path: '/setting/privilege', component: PrivList, name: '权限列表' },
        // { path: '/setting/demo', component: Demo, name: 'demo' }
      ]
    },
    // 彩票管理
    {
      path: '/',
      component: Home,
      name: '彩票管理',
      iconCls: 'el-icon-menu',
      children: [
        {path: '/lottery/open/lottery',                         component: require('@/components/lottery/winning/list'),           name: '开奖信息'},
        {path: '/lottery/open/history/:lottery',                component: require('@/components/lottery/winning/history'),        name: '历史记录',       hidden:true},
        {path: '/lottery/buycai/options',                       component: require('@/components/lottery/BuycaiOptions'),          name: '购彩配置'},
        {path: '/lottery/buycai/playtimesetting',               component: require('@/components/lottery/PlayTimeSettings'),       name: '玩法时间配置'},
        {path: '/lottery/buycai/options/add',                   component: require('@/components/lottery/BuycaiOptionsAdd'),       name: '添加购彩配置',    hidden:true},
        {path: '/lottery/buycai/options/edit/:id/lottery/:lid', component: require('@/components/lottery/BuycaiOptionsEdit'),      name: '编辑购彩配置',    hidden:true},
        {path: '/lottery/home/options',                         component: require('@/components/lottery/lotteryHomeOptions'),     name: '首页彩种配置'},
        {path: '/lottery/home/options/add',                     component: require('@/components/lottery/lotteryHomeOptionsAdd'),  name: '新增首页彩种配置', hidden:true},
        {path: '/lottery/home/options/edit/:id',                component: require('@/components/lottery/lotteryHomeOptionsEdit'), name: '编辑首页彩种配置', hidden:true},
        // {path: '/lottery/home/options/winning/list',  component: require('@/components/winning/list'),                   name: '首页最新中奖'},
        ]
    },
      {
        path : "/",
        component: Home,
        name: '活动管理',
        iconCls: 'el-icon-menu',
        children: [
            { path: '/activity/cdkey/list',       component: require('@/components/activity/cdkey/list'),        name: '兑换码' },
            { path: '/activity/cdkey/add',        component: require('@/components/activity/cdkey/add'),         name: '新增兑换码', hidden:true },
            { path: '/activity/cdkey/detail/:id', component: require('@/components/activity/cdkey/detail'),      name: '兑换码详细', hidden:true },
            { path: '/activity/cdkey/edit/:id',   component: require('@/components/activity/cdkey/edit'),        name: '编辑兑换码', hidden:true },
            { path: '/activity/task/taskList',   component: require('@/components/activity/taskList'),           name: '任务管理' },
            { path: '/activity/activity/activityList',   component: require('@/components/activity/activityList'),          name: '活动' },
            // { path: '/activity/task/taskList',   component: require('@/components/activity/taskList'),          name: '任务管理' },
            { path: '/activity/task/add',   component: require('@/components/activity/addTask'),          name: '添加任务',hidden:true },
            { path: '/activity/task/edit/:id',   component: require('@/components/activity/editTask'),          name: '编辑',hidden:true },
            { path: '/activity/activity/add',   component: require('@/components/activity/addActivity'),          name: '添加活动',hidden:true },
            { path: '/activity/activity/edit/:id',   component: require('@/components/activity/editActivity'),          name: '编辑活动',hidden:true },
        ]
      },
        {
        path : "/",
        component: Home,
        name: '礼包模板管理',
        iconCls: 'el-icon-menu',
        children: [
            { path: '/gift/template/list',       component: require('@/components/gift/template/list'),   name: '礼包模板' },
            { path: '/gift/template/add',        component: require('@/components/gift/template/add'),    name: '新增礼包模板',hidden:true },
            { path: '/gift/template/detail/:id', component: require('@/components/gift/template/detail'), name: '礼包模板详细',hidden:true },
            { path: '/gift/template/edit/:id',   component: require('@/components/gift/template/edit'),   name: '编辑礼包模板',hidden:true },
        ]
      },
    {
      path : "/",
      component: Home,
      name: '平台配置',
      iconCls: 'el-icon-menu',
      children: [
        { path: '/options/contact/index',       component: require('@/components/options/contact/index'),   name: '客服信息' },
        { path: '/options/feedback/list',       component: require('@/components/options/feedback/list'),   name: '用户反馈' },
        { path: '/options/feedback/detail/:id', component: require('@/components/options/feedback/detail'), name: '处理反馈', hidden:true },
        { path: '/options/faq/list', component: require('@/components/options/faqs/FAQList'), name: '常见问题' },
        { path: '/options/faq/add', component: require('@/components/options/faqs/FAQForm'), name: '添加常见问题', hidden:true },
        { path: '/options/faq/edit/:id', component: require('@/components/options/faqs/FAQForm'), name: '修改常见问题', hidden:true },
      ]
    },
    {
      path : "/",
      component: Home,
      name: '用户管理',
      iconCls: 'el-icon-menu',
      children: [
        { path: '/usercenter/user/list',       component: require('@/components/usercenter/user/list'),   name: '用户列表'},
        { path: '/usercenter/user/detail/:id', component: require('@/components/usercenter/user/detail'), name: '用户详细信息', hidden:true},
        { path: '/usercenter/user/fund/:id', component: require('@/components/usercenter/user/fund'), name: '资金详细', hidden: true },
        { path: '/usercenter/withdraw/list', component: require('@/components/withdraw/WithdrawList'), name: '提现审核'},
        { path: '/usercenter/withdraw/audit/:id', component: require('@/components/withdraw/WithdrawAudit'), name: '提现审核详情', hidden: true},
        { path: '/usercenter/withdraw/transfer', component: require('@/components/withdraw/WithdrawTransferList'), name: '提现转账列表' },
        { path: '/usercenter/withdraw/transfer/detail/:id', component: require('@/components/withdraw/WithdrawTransferForm'), name: '确认转账', hidden:true },
        { path: '/usercenter/withdraw/auth', component: require('@/components/withdraw/WithdrawAuthList'), name: '提现操作权限' },
        { path: '/usercenter/withdraw/auth/add', component: require('@/components/withdraw/WithdrawAuthForm'), name: '添加提现申请权限', hidden: true },
        { path: '/usercenter/withdraw/auth/edit/:id', component: require('@/components/withdraw/WithdrawAuthForm'), name: '编辑提现权限', hidden: true }
      ]
    },

    //订单管理.
    {
      path : "/",
      component: Home,
      name: '订单管理',
      iconCls: 'el-icon-menu',
      children: [
          { path: '/order', component: Order, name: '订单列表' },
          { path: '/order/user', component: require('@/components/order/userOrder'), name: '用户订单' },
          { path: '/order/buycai', component: require('@/components/order/buycaiOrder'), name: '购彩订单' },
      ]
    },
    //日志管理.
    {
      path : "/",
      component:Home,
      name: "日志管理",
      iconCls: 'el-icon-menu',
      children: [
          { path: '/log/list', component: require('@/components/log/OperatingLog'), name: '操作日志' },
      ]
    },
    {
      path : "/",
      component:Home,
      name: "发现管理",
      iconCls: 'el-icon-menu',
      children: [
          { path: '/news', component: News, name: '新闻管理' },
          {
            path: '/news/add',
            component: NewsForm,
            name: '添加新闻',
            hidden: true
          },
          {
                name: '修改新闻',
                path: '/news/edit/:id',
                component: NewsForm,
                hidden: true,
          },
          {
            name: 'banner图管理',
            path: '/discover/banner',
            component: DiscoverBanner,
          },
          {
            name: '添加banner',
            path: '/discover/banner/add',
            component: require('@/components/discover/BannerForm'),
            hidden: true
          },
          {
            name: '修改banner',
            path: '/discover/banner/edit/:id',
            component: require('@/components/discover/BannerForm'),
            hidden: true
          },
      ]
    },
    //足彩管理
    {
      path : "/",
      component:Home,
      name: "竞彩管理",
      iconCls: 'el-icon-menu',
      children: [
        {
          name: '比赛管理',
          path: '/football/game',
          component: require('@/components/football/GameList'),
        },
        {
          name: '添加比赛',
          path: '/football/game/add',
          component: require('@/components/football/GameForm'),
          hidden: true
        },
        {
          name: '修改比赛',
          path: '/football/game/edit/:id',
          component: require('@/components/football/GameForm'),
          hidden: true,
        },
         {
          name: '查看赔率',
          path: '/football/odds/:id',
          component: require('@/components/football/OddsList'),
          hidden: true
        },
        {
          name: '添加&修改赔率',
          path: '/football/odds/:op/:id',
          component: require('@/components/football/OddsForm'),
          hidden: true
        },
        {
          name: '开奖管理',
          path: '/football/opencai',
          component: require('@/components/football/Opencai'),
        },
        {
          name: '修改开奖',
          path: '/football/opencai/edit/:id',
          component: require('@/components/football/OpencaiForm'),
          hidden:true
        },
        {
          name: '球队管理',
          path: '/football/team',
          component: require('@/components/football/TeamList'),
        },
        {
          name: '添加球队',
          path: '/football/team/add',
          component: require('@/components/football/TeamForm'),
          hidden:true
        },
        {
          name: '修改球队信息',
          path: '/football/team/edit/:id',
          component: require('@/components/football/TeamForm'),
          hidden: true,
        },
        {
          name: '联赛管理',
          path: '/football/league',
          component: require('@/components/football/LeagueList'),
        },
      ]
    },
    // 跳转至404
    {
      path: '*',
      hidden: true,
      redirect: { path: '/404' }
    }
  ]
})

