(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-32ab31e2"],{"0844":function(e,r,t){},4565:function(e,r,t){"use strict";var s=t("b195"),a=t.n(s);a.a},"489d":function(e,r,t){"use strict";var s=t("0844"),a=t.n(s);a.a},b195:function(e,r,t){},d5c2:function(e,r,t){"use strict";t.r(r);var s=function(){var e=this,r=e.$createElement,t=e._self._c||r;return t("div",{staticStyle:{"background-color":"#3d3744"}},[t("div",{staticClass:"video-container"},[t("video",{attrs:{src:"https://static1.gotokeep.com/homepage/full.mp4",autoplay:"",loop:"",muted:"",id:"video-bg"},domProps:{muted:!0}})]),e._v(" "),t("div",{staticClass:"login-container"},[t("el-form",{ref:"ruleForm",staticClass:"login-form",attrs:{model:e.ruleForm,rules:e.rules,"auto-complete":"on","label-position":"left"}},[t("div",{staticClass:"title-container"},[t("h3",{staticClass:"title"},[e._v("美术世界拼团-管理用户注册V0.1")])]),e._v(" "),t("el-form-item",{attrs:{prop:"username"}},[t("span",{staticClass:"svg-container"},[t("svg-icon",{attrs:{"icon-class":"user"}})],1),e._v(" "),t("el-input",{attrs:{placeholder:"请输入用户名","auto-complete":"on"},model:{value:e.ruleForm.username,callback:function(r){e.$set(e.ruleForm,"username",r)},expression:"ruleForm.username"}})],1),e._v(" "),t("el-form-item",{attrs:{prop:"pass"}},[t("span",{staticClass:"svg-container"},[t("svg-icon",{attrs:{"icon-class":"password"}})],1),e._v(" "),t("el-input",{attrs:{placeholder:"请输入最少7位密码",autocomplete:"off"},model:{value:e.ruleForm.pass,callback:function(r){e.$set(e.ruleForm,"pass",r)},expression:"ruleForm.pass"}})],1),e._v(" "),t("el-form-item",{attrs:{prop:"checkPass"}},[t("span",{staticClass:"svg-container"},[t("svg-icon",{attrs:{"icon-class":"password"}})],1),e._v(" "),t("el-input",{attrs:{placeholder:"请重新输入密码",autocomplete:"off"},model:{value:e.ruleForm.checkPass,callback:function(r){e.$set(e.ruleForm,"checkPass",r)},expression:"ruleForm.checkPass"}})],1),e._v(" "),t("el-form-item",{attrs:{prop:"code"}},[t("span",{staticClass:"svg-container"},[t("svg-icon",{attrs:{"icon-class":"password"}})],1),e._v(" "),t("el-input",{attrs:{placeholder:"请输入邀请码 pinpin","auto-complete":"on"},model:{value:e.ruleForm.code,callback:function(r){e.$set(e.ruleForm,"code",r)},expression:"ruleForm.code"}})],1),e._v(" "),t("el-button",{staticStyle:{width:"100%","margin-bottom":"30px","background-color":"#24c789",border:"0"},attrs:{type:"primary"},nativeOn:{click:function(r){return r.preventDefault(),e.handleLogin(r)}}},[e._v("注册")])],1)],1)])},a=[],o=(t("96cf"),t("3b8d")),n=t("c24f"),c=(t("bc3a"),{name:"Login",data:function(){var e=this,r=function(e,r,t){if(!r)return t(new Error("用户名不能为空"));t()},t=function(r,t,s){""===t||t.length<=6?s(new Error("请输入密码或者检查密码位数最少7位")):(""!==e.ruleForm.checkPass&&e.$refs.ruleForm.validateField("checkPass"),s())},s=function(r,t,s){""===t||t.length<=6?s(new Error("请再次输入密码或者检查密码位数最少7位")):t!==e.ruleForm.pass?s(new Error("两次输入密码不一致!")):s()},a=function(e,r,t){""===r?t(new Error("请输入注册验证码")):"pinpin"!==r?t(new Error("请输入正确的验证码!")):t()};return{ruleForm:{username:"",pass:"",checkPass:"",code:""},rules:{username:[{trigger:"blur",validator:r}],pass:[{validator:t,trigger:"blur"}],checkPass:[{validator:s,trigger:"blur"}],code:[{validator:a,trigger:"blur"}]},passwordType:"password",redirect:void 0}},watch:{$route:{handler:function(e){this.redirect=e.query&&e.query.redirect},immediate:!0}},methods:{handleLogin:function(){var e=this;this.$refs.ruleForm.validate(function(){var r=Object(o["a"])(regeneratorRuntime.mark((function r(t){var s,a;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:if(!t){r.next=13;break}return s={name:e.ruleForm.username,passWord:e.ruleForm.pass},r.next=4,Object(n["c"])(s).catch((function(e){return e}));case 4:if(a=r.sent,console.log(a,"res-------res-----"),20001!==a.code){r.next=8;break}return r.abrupt("return",e.$message.error(a.msg));case 8:if(2e4===a.code){r.next=10;break}return r.abrupt("return",e.$message.error(a.msg));case 10:e.$router.push({path:e.redirect||"/login"}),r.next=15;break;case 13:return console.log("error submit!!"),r.abrupt("return",!1);case 15:case"end":return r.stop()}}),r)})));return function(e){return r.apply(this,arguments)}}())}}}),i=c,l=(t("4565"),t("489d"),t("2877")),u=Object(l["a"])(i,s,a,!1,null,"48dfe7f4",null);r["default"]=u.exports}}]);