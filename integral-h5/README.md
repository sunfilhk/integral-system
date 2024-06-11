## 项目简介
```
    本项目是使用uniapp所写的h5手机端，主要有首页、客户信息的增删改查、客户投资的增删改查、客户信息和投资的审核和审核记录的
展示、我的页面展示用户个人信息和退出登录、新增下级业务员的审核和审核记录的展示、展示下级业务员的功能
```


## 目录结构和功能

```
zhongyue ：主目录
  common ：公共工具js
    |-- api.js ：后台接口
    |-- costum.js ：方法集合
    |-- request.js ：请求配置
  pages ：页面组件
    |-- index/index.vue ：首页
    |-- customer/customer.vue : 业务：展示客户信息、新增客户和客户投资列表
    |-- incustomer/incustomer.vue : 业务处理：添加或者修改客户信息和投资
    |-- infor/infor.vue : 客户管理：展示详细的客户信息和修改信息
    |-- investment/investment.vue : 客户投资：展示客户投资信息、新增、修改、撤销、终止投资
    |-- ininvestment/ininvestment.vue : 投资处理：添加或者修改投资
    |-- examines/examines.vue : 审核：审核客户的添加和修改以及客户投资的添加、修改和终止
    |-- record/record.vue : 审核记录：审核记录的展示
    |-- inexamine/inexamine.vue : 审核管理：展示该条审核的详细信息以及审核同意或者驳回
    |-- mine/mine.vue : 我的：展示用户的个人信息、退出登录、下级业务员审核、下级业务员、业务员审核记录、修改密码的导向操作
    |-- salesman/salesman.vue : 用户审核：新增下级业务员和审核
    |-- level/level.vue : 新增用户：填写新增下级业务员信息
    |-- subordinate/subordinate.vue : 下级业务员：展示下级业务员信息
    |-- audituser/audituser.vue : 审核记录：下级业务员的审核记录
    |-- pass/pass.vue : 修改密码：修改用户密码
    |-- login/login.vue : 登录
  static ：静态资源
  uni_modules ：依赖
  unpackage ：编译文件
  App.vue ：应用配置，用来配置App全局样式以及监听 应用生命周期
  main.js ：Vue初始化入口文件
  manifest.json ：配置应用名称、appid、logo、版本等打包信息
  pages.json ：配置页面路由、导航条、选项卡等页面类信息
```


## 外部插件
   表单插件:/uni_modules/hpy-form-select
