(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-31430a74"],{4692:function(t,a,e){},"626e":function(t,a,e){"use strict";e.d(a,"i",(function(){return r})),e.d(a,"e",(function(){return i})),e.d(a,"h",(function(){return s})),e.d(a,"f",(function(){return o})),e.d(a,"g",(function(){return c})),e.d(a,"l",(function(){return u})),e.d(a,"j",(function(){return l})),e.d(a,"k",(function(){return p})),e.d(a,"a",(function(){return d})),e.d(a,"d",(function(){return f})),e.d(a,"b",(function(){return m})),e.d(a,"c",(function(){return v}));var n=e("b775");function r(t){return Object(n["a"])({url:"/api/v1/profitconfigOnce",method:"post",params:t})}function i(t){return Object(n["a"])({url:"/api/v1/getonceinto",method:"get",params:t})}function s(t){return Object(n["a"])({url:"/api/v1/profitconfigList",method:"get",params:t})}function o(t){return Object(n["a"])({url:"/api/v1/getProfitcustomers",method:"get",params:t})}function c(t){return Object(n["a"])({url:"/api/v1/profitconfigEdit",method:"post",data:t})}function u(t){return Object(n["a"])({url:"/api/v1/userList",method:"get",params:t})}function l(t){return Object(n["a"])({url:"/api/v1/teamList",method:"get",params:t})}function p(t){return Object(n["a"])({url:"/api/v1/updateProfitConfigOnce",method:"post",params:t})}function d(t){return Object(n["a"])({url:"/api/v1/delProfitconfigOnce",method:"post",params:t})}function f(t){return Object(n["a"])({url:"/api/v1/getUserVipLevel",method:"get",params:t})}function m(t){return Object(n["a"])({url:"/api/v1/editUserVipLevel",method:"post",params:t})}function v(t){return Object(n["a"])({url:"/api/v1/editlevellock",method:"post",params:t})}},9406:function(t,a,e){"use strict";e.r(a);var n=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",{staticStyle:{padding:"20px"}},[e("el-card",{staticClass:"topCard"},[e("div",{staticClass:"data"},[e("el-row",[e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("用户姓名："+t._s(t.data.nick_name))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("历史总业绩: "+t._s(t.data.total_performance))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("当月业绩: "+t._s(t.data.month))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("当日业绩: "+t._s(t.data.today_performance))])]),e("el-col",{attrs:{span:24}},[e("div",{staticClass:"a"},[t._v("vip等级: "+t._s(t.data.vip_level))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("vip称号: "+t._s(t.data.vip_title))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("历史总利润: "+t._s(t.profit.total))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("当月利润: "+t._s(t.profit.month))])]),e("el-col",{attrs:{span:6}},[e("div",{staticClass:"a"},[t._v("当日利润: "+t._s(t.profit.today))])])],1)],1)]),e("el-card",[e("div",{staticStyle:{margin:"20px 0 20px","font-weight":"700"}},[t._v("下级业务员")]),e("el-table",{staticStyle:{"min-width":"800px"},attrs:{border:"",data:t.list,"header-cell-style":{textAlign:"center"},"cell-style":{textAlign:"center"}}},[e("el-table-column",{attrs:{sortable:"",label:"等级",prop:"vip",width:"100"}}),e("el-table-column",{attrs:{label:"姓名",prop:"nick_name"}}),e("el-table-column",{attrs:{label:"用户名",prop:"username"}}),e("el-table-column",{attrs:{label:"手机号",prop:"phone"}}),e("el-table-column",{attrs:{sortable:"",label:"团队业绩",prop:"performance"}})],1),e("pagination",{directives:[{name:"show",rawName:"v-show",value:t.total>0,expression:"total > 0"}],attrs:{total:t.total,page:t.queryParams.pageIndex,limit:t.queryParams.pageSize},on:{"update:page":function(a){return t.$set(t.queryParams,"pageIndex",a)},"update:limit":function(a){return t.$set(t.queryParams,"pageSize",a)},pagination:t.auto}})],1)],1)},r=[],i=e("c7eb"),s=(e("96cf"),e("1da1")),o=e("5530"),c=e("2f62"),u=e("c24f"),l=e("626e"),p={name:"Dashboard",data:function(){return{data:{},deptId:"",profit:{},list:[],queryParams:{pageIndex:1,pageSize:10},total:0,arr:[]}},computed:Object(o["a"])({},Object(c["b"])(["roles"])),created:function(){this.auto(),this.getData()},methods:{getData:function(){var t=Object(s["a"])(Object(i["a"])().mark((function t(){var a;return Object(i["a"])().wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(u["c"])();case 2:a=t.sent,this.profit=a.data;case 4:case"end":return t.stop()}}),t,this)})));function a(){return t.apply(this,arguments)}return a}(),auto:function(){var t=Object(s["a"])(Object(i["a"])().mark((function t(){var a=this;return Object(i["a"])().wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(u["b"])().then((function(t){a.data=t.data,a.deptId=t.data.deptId}));case 2:Object(l["h"])({profittype:"2",deptId:this.deptId}).then((function(t){a.arr=t.data.map((function(t){return t.profitlevel})),Object(u["d"])(a.queryParams).then((function(t){a.total=t.total,a.list=t.list}))}));case 3:case"end":return t.stop()}}),t,this)})));function a(){return t.apply(this,arguments)}return a}()}},d=p,f=(e("d920"),e("2877")),m=Object(f["a"])(d,n,r,!1,null,"7b8728f2",null);a["default"]=m.exports},d920:function(t,a,e){"use strict";e("4692")}}]);