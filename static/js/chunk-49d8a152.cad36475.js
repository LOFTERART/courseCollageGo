(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-49d8a152"],{2109:function(e,t,r){},"787a":function(e,t,r){"use strict";var a=r("2109"),n=r.n(a);n.a},8325:function(e,t,r){"use strict";var a=r("b311");if(!a)throw new Error("you should npm install `clipboard` --save at first ");var n={bind:function(e,t){if("success"===t.arg)e._v_clipboard_success=t.value;else if("error"===t.arg)e._v_clipboard_error=t.value;else{var r=new a(e,{text:function(){return t.value},action:function(){return"cut"===t.arg?"cut":"copy"}});r.on("success",(function(t){var r=e._v_clipboard_success;r&&r(t)})),r.on("error",(function(t){var r=e._v_clipboard_error;r&&r(t)})),e._v_clipboard=r}},update:function(e,t){"success"===t.arg?e._v_clipboard_success=t.value:"error"===t.arg?e._v_clipboard_error=t.value:(e._v_clipboard.text=function(){return t.value},e._v_clipboard.action=function(){return"cut"===t.arg?"cut":"copy"})},unbind:function(e,t){"success"===t.arg?delete e._v_clipboard_success:"error"===t.arg?delete e._v_clipboard_error:(e._v_clipboard.destroy(),delete e._v_clipboard)}},o=function(e){e.directive("Clipboard",n)};window.Vue&&(window.clipboard=n,Vue.use(o)),n.install=o;t["a"]=n},a647:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("el-table",{staticStyle:{width:"100%"},attrs:{stripe:"",data:e.tableData}},[r("el-table-column",{attrs:{label:"#",type:"index"}}),e._v(" "),r("el-table-column",{attrs:{label:"订单创建时间",prop:"created_at"}}),e._v(" "),r("el-table-column",{attrs:{label:"姓名",prop:"address.true_name"}}),e._v(" "),r("el-table-column",{attrs:{label:"电话",prop:"address.tel_phone"}}),e._v(" "),r("el-table-column",{attrs:{label:"省份",prop:"address.province"}}),e._v(" "),r("el-table-column",{attrs:{label:"详细地址",width:"400"},scopedSlots:e._u([{key:"default",fn:function(t){var a=t.row;return[r("span",[e._v(e._s(a.address.full_address_add))]),e._v(" "),r("el-button",{directives:[{name:"clipboard",rawName:"v-clipboard:copy",value:a.address.full_address_add,expression:"row.address.full_address_add",arg:"copy"},{name:"clipboard",rawName:"v-clipboard:success",value:e.clipboardSuccess,expression:"clipboardSuccess",arg:"success"}],attrs:{type:"text",size:"mini"}},[e._v("\n          复制\n        ")])]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"科目",prop:"course.name"}}),e._v(" "),r("el-table-column",{attrs:{label:"课程内容",prop:"course.courseName"}}),e._v(" "),r("el-table-column",{attrs:{align:"right"},scopedSlots:e._u([{key:"default",fn:function(t){return[r("el-button",{attrs:{slot:"reference",icon:"el-icon-truck",size:"mini",type:"danger"},on:{click:function(r){return e.handleSendOK(t.$index,t.row)}},slot:"reference"},[e._v(e._s(1===t.row.send_type?"发货":"已发货"))])]}}])})],1),e._v(" "),r("div",{staticStyle:{"margin-top":"20px",float:"right"}},[r("el-pagination",{attrs:{layout:"total,sizes,prev, pager, next","page-sizes":[10,15,20],total:e.total},on:{"size-change":e.handleSizeChange,"current-change":e.handleCurrentChange}})],1),e._v(" "),r("el-dialog",{attrs:{title:"发货信息",visible:e.dialogFormVisible},on:{"update:visible":function(t){e.dialogFormVisible=t}}},[r("el-form",{ref:"ruleForm",attrs:{model:e.form,rules:e.rules}},[r("el-form-item",{attrs:{label:"快递单号","label-width":e.formLabelWidth,prop:"exNum"}},[r("el-input",{attrs:{autocomplete:"off"},model:{value:e.form.exNum,callback:function(t){e.$set(e.form,"exNum",t)},expression:"form.exNum"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"快递公司","label-width":e.formLabelWidth}},[r("el-select",{attrs:{placeholder:"请选择快递公司"},model:{value:e.value,callback:function(t){e.value=t},expression:"value"}},e._l(e.form.options,(function(e){return r("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1)],1),e._v(" "),r("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogFormVisible=!1}}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:e.clickOk}},[e._v("确 定")])],1)],1)],1)},n=[],o=(r("c5f6"),r("7f7f"),r("96cf"),r("3b8d")),l=r("db72"),s=r("f8b7"),i=r("8325"),c=r("2f62"),u={name:"index",directives:{clipboard:i["a"]},computed:Object(l["a"])({},Object(c["b"])(["name"])),data:function(){return{page:1,size:10,total:0,tableData:[],value:"",dialogFormVisible:!1,form:{exNum:"",orderID:null,options:[{value:"申通快递",label:"申通快递"},{value:"圆通快递",label:"圆通快递"},{value:"中通快递",label:"中通快递"},{value:"韵达快递",label:"韵达快递"},{value:"顺丰快递",label:"顺丰快递"}]},rules:{exNum:[{required:!0,message:"请输入快递单号",trigger:"blur"},{min:5,max:30,message:"长度在 5 到 30 个字符",trigger:"blur"}]},formLabelWidth:"80px"}},created:function(){this.fetchData(this.page,this.size)},methods:{handleSizeChange:function(e){console.log("每页 ".concat(e," 条")),this.size=e,this.fetchData(this.page,this.size),console.log(11,"-----每页----")},handleCurrentChange:function(e){console.log("当前页: ".concat(e)),this.page=e,this.fetchData(e,this.size),console.log(11,"-----当前页----")},fetchData:function(){var e=Object(o["a"])(regeneratorRuntime.mark((function e(t,r){var a,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return a={send_type:1,page:t,size:r},e.next=3,Object(s["a"])(a).catch((function(e){return e}));case 3:if(n=e.sent,2e4===n.code){e.next=6;break}return e.abrupt("return",this.$message.error(n.msg));case 6:this.tableData=n.data.items,this.total=n.data.total;case 8:case"end":return e.stop()}}),e,this)})));function t(t,r){return e.apply(this,arguments)}return t}(),handleSendOK:function(e,t){this.dialogFormVisible=!0,this.form.orderID=t.id},clickOk:function(){var e=this;this.$refs.ruleForm.validate(function(){var t=Object(o["a"])(regeneratorRuntime.mark((function t(r){var a,n;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(!r){t.next=14;break}return a={operator:e.name,send_type:2,ExName:e.value,exRegion:e.form.exNum,order_id:Number(e.form.orderID)},t.next=4,Object(s["b"])(a).catch((function(e){return e}));case 4:if(n=t.sent,console.log(n,"-----res-----"),2e4===n.code){t.next=8;break}return t.abrupt("return",e.$message.error(n.msg));case 8:e.dialogFormVisible=!1,e.$refs.ruleForm.resetFields(),e.fetchData(e.page,e.size),e.$message({message:"恭喜你，发货成功",type:"success"}),t.next=16;break;case 14:return console.log("error submit!!"),t.abrupt("return",!1);case 16:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}())},clipboardSuccess:function(){this.$message({message:"复制成功",type:"success",duration:1500})}}},d=u,p=(r("787a"),r("2877")),b=Object(p["a"])(d,a,n,!1,null,"47543ab8",null);t["default"]=b.exports},f8b7:function(e,t,r){"use strict";r.d(t,"a",(function(){return n})),r.d(t,"b",(function(){return o}));var a=r("b775");function n(e){return Object(a["a"])({url:"/api/v1/getOrders",method:"post",data:e})}function o(e){return Object(a["a"])({url:"/api/v1/order/upSendType",method:"post",data:e})}}}]);