(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-05644fd6"],{"6a84":function(e,t,r){},aa77:function(e,t,r){var n=r("5ca1"),a=r("be13"),o=r("79e5"),i=r("fdef"),s="["+i+"]",c="​",l=RegExp("^"+s+s+"*"),u=RegExp(s+s+"*$"),p=function(e,t,r){var a={},s=o((function(){return!!i[e]()||c[e]()!=c})),l=a[e]=s?t(m):i[e];r&&(a[r]=l),n(n.P+n.F*s,"String",a)},m=p.trim=function(e,t){return e=String(a(e)),1&t&&(e=e.replace(l,"")),2&t&&(e=e.replace(u,"")),e};e.exports=p},b0b6:function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"app-container"},[r("el-select",{attrs:{placeholder:"分类筛选",clearable:""},model:{value:e.value,callback:function(t){e.value=t},expression:"value"}},e._l(e.options,(function(e){return r("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1),e._v(" "),r("el-button",{staticStyle:{"margin-left":"20px","margin-right":"20px"},attrs:{type:"primary",icon:"el-icon-search"},on:{click:e.clickSearch}},[e._v("搜索")]),e._v(" "),r("el-button",{staticClass:"addbutton",attrs:{type:"success",icon:"el-icon-plus"},on:{click:function(t){e.dialogVisible=!0}}},[e._v("添加书籍")]),e._v(" "),r("el-table",{attrs:{data:e.books,stripe:""}},[r("el-table-column",{attrs:{type:"index",label:"#"}}),e._v(" "),r("el-table-column",{attrs:{prop:"course.name",sortable:"",label:"课程"}}),e._v(" "),r("el-table-column",{attrs:{prop:"title",label:"书籍"}}),e._v(" "),r("el-table-column",{attrs:{prop:"tip",label:"标记"}}),e._v(" "),r("el-table-column",{attrs:{prop:"create_people",label:"创建人"}}),e._v(" "),r("el-table-column",{attrs:{label:"图片"},scopedSlots:e._u([{key:"default",fn:function(e){return[r("el-image",{staticStyle:{width:"40px",height:"40px"},attrs:{src:e.row.img_url,"preview-src-list":[e.row.img_url]}})]}}])}),e._v(" "),r("el-table-column",{attrs:{label:"操作"},scopedSlots:e._u([{key:"default",fn:function(t){var n=t.row;return[r("el-button",{attrs:{type:"text"},on:{click:function(t){return e.clickDelBook(n)}}},[e._v("删除")])]}}])})],1),e._v(" "),r("el-dialog",{attrs:{"show-close":!1,title:"增加分类",visible:e.dialogVisible,width:"50%"},on:{"update:visible":function(t){e.dialogVisible=t}}},[r("el-form",{ref:"ruleForm",staticClass:"demo-ruleForm",attrs:{model:e.ruleForm,rules:e.rules,"label-width":"100px"}},[r("el-form-item",{attrs:{label:"书籍名字",prop:"name"}},[r("el-input",{model:{value:e.ruleForm.name,callback:function(t){e.$set(e.ruleForm,"name",t)},expression:"ruleForm.name"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"书籍提示",prop:"tip"}},[r("el-input",{model:{value:e.ruleForm.tip,callback:function(t){e.$set(e.ruleForm,"tip",t)},expression:"ruleForm.tip"}})],1),e._v(" "),r("el-form-item",{attrs:{label:"书籍图片",prop:"img"}},[r("el-upload",{staticClass:"upload-demo",attrs:{name:"upload[]",drag:"",limit:1,"auto-upload":"","show-file-list":"","file-list":e.fileList,action:e.UPLOADURL,"on-exceed":e.clickImgExceed,"on-preview":e.handlePreview,"on-remove":e.handleRemove,"before-upload":e.beforeAvatarUpload,"on-success":e.handleAvatarSuccess}},[r("i",{staticClass:"el-icon-upload"}),e._v(" "),r("div",{staticClass:"el-upload__text"},[e._v("\n            将文件拖到此处，或"),r("em",[e._v("点击上传")]),e._v(" "),r("div",{staticStyle:{color:"#ccc","font-size":"12px"}},[e._v("只能上传"),r("em",[e._v("1张")]),e._v(" jpg文件，且不超过1M,分辨率为300*300")])])])],1),e._v(" "),r("el-form-item",{attrs:{label:"课程目录",prop:"options"}},[r("el-select",{attrs:{placeholder:"请选择课程目录"},model:{value:e.courseId,callback:function(t){e.courseId=t},expression:"courseId"}},e._l(e.ruleForm.options,(function(e){return r("el-option",{key:e.id,attrs:{label:e.name,value:e.id}})})),1)],1)],1),e._v(" "),r("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[r("el-button",{on:{click:function(t){e.dialogVisible=!1}}},[e._v("取 消")]),e._v(" "),r("el-button",{attrs:{type:"primary"},on:{click:function(t){return e.clickDone("ruleForm")}}},[e._v("确 定")])],1)],1)],1)},a=[],o=(r("7f7f"),r("c5f6"),r("96cf"),r("3b8d")),i=r("db72"),s=r("cabd"),c=r("2f62"),l={name:"index",computed:Object(i["a"])({},Object(c["b"])(["name"])),data:function(){return{books:[],UPLOADURL:s["a"],imageUrl:"",fileList:[],dialogVisible:!1,ruleForm:{name:"",tip:"",img:"",options:[]},courseId:null,rules:{name:[{required:!0,message:"请输入书籍名称",trigger:"blur"},{min:3,max:20,message:"长度在 3 到 20 个字符",trigger:"blur"}],tip:[{required:!1,message:"请填写书籍运营标记",trigger:"blur"}],img:[{required:!0,message:"请上传图片",trigger:"blur"}],options:[{required:!0,message:"请选择关联课程目录",trigger:"blur"}]},options:[],value:""}},created:function(){this.getCourse(),this.getBooks({course_id:0})},methods:{clickDelBook:function(e){var t=this;console.log(e),this.$confirm("此操作将删除该书籍, 是否继续?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(Object(o["a"])(regeneratorRuntime.mark((function r(){var n,a;return regeneratorRuntime.wrap((function(r){while(1)switch(r.prev=r.next){case 0:return n={id:Number(e.id)},r.next=3,Object(s["e"])(n).catch((function(e){return e}));case 3:if(a=r.sent,console.log(a,"shanchuzhihou"),2e4===a.code){r.next=7;break}return r.abrupt("return",t.$message.error(a.msg));case 7:t.getBooks(),t.$message({type:"success",message:a.msg});case 9:case"end":return r.stop()}}),r)})))).catch((function(){t.$message({type:"info",message:"已取消删除"})}))},getCourse:function(){var e=Object(o["a"])(regeneratorRuntime.mark((function e(){var t;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Object(s["h"])().catch((function(e){return e}));case 2:if(t=e.sent,2e4===t.code){e.next=5;break}return e.abrupt("return",this.$message.error(t.msg));case 5:this.ruleForm.options=t.data,this.options=t.data;case 7:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),getBooks:function(){var e=Object(o["a"])(regeneratorRuntime.mark((function e(t){var r;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return e.next=2,Object(s["f"])(t).catch((function(e){return e}));case 2:if(r=e.sent,2e4===r.code){e.next=5;break}return e.abrupt("return",this.$message.error(r.msg));case 5:this.books=r.data;case 6:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),clickDone:function(e){var t=this,r=this.name;this.$refs[e].validate(function(){var e=Object(o["a"])(regeneratorRuntime.mark((function e(n){var a,o;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(!n){e.next=15;break}return a={create_people:r,title:t.ruleForm.name,tip:t.ruleForm.tip,img:t.imageUrl,course_id:Number(t.courseId)},console.log(a,"pass params"),e.next=5,Object(s["c"])(a).catch((function(e){return e}));case 5:if(o=e.sent,2e4===o.code){e.next=8;break}return e.abrupt("return",t.$message.error(o.msg));case 8:t.dialogVisible=!1,t.imageUrl="",t.fileList=[],t.$refs.ruleForm.resetFields(),t.getBooks(),e.next=16;break;case 15:return e.abrupt("return",!1);case 16:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}())},clickSearch:function(){var e={course_id:this.value||0};this.getBooks(e)},clickImgExceed:function(e,t){console.log(e,t)},handleRemove:function(e,t){console.log(e,t)},handlePreview:function(e){console.log(e)},handleAvatarSuccess:function(e,t){this.imageUrl=e.files.imgName,this.ruleForm.img=e.files.imgName,console.log(this.imageUrl,"-----",e)},beforeAvatarUpload:function(e){var t="image/jpeg"===e.type,r=e.size/1024/1024<1;return t||this.$message.error("上传图片只能是 JPG 格式!"),r||this.$message.error("上传图片大小不能超过 1MB!"),t&&r}}},u=l,p=(r("e659"),r("2877")),m=Object(p["a"])(u,n,a,!1,null,"4c9fd28e",null);t["default"]=m.exports},c5f6:function(e,t,r){"use strict";var n=r("7726"),a=r("69a8"),o=r("2d95"),i=r("5dbc"),s=r("6a99"),c=r("79e5"),l=r("9093").f,u=r("11e9").f,p=r("86cc").f,m=r("aa77").trim,f="Number",d=n[f],b=d,g=d.prototype,v=o(r("2aeb")(g))==f,h="trim"in String.prototype,k=function(e){var t=s(e,!1);if("string"==typeof t&&t.length>2){t=h?t.trim():m(t,3);var r,n,a,o=t.charCodeAt(0);if(43===o||45===o){if(r=t.charCodeAt(2),88===r||120===r)return NaN}else if(48===o){switch(t.charCodeAt(1)){case 66:case 98:n=2,a=49;break;case 79:case 111:n=8,a=55;break;default:return+t}for(var i,c=t.slice(2),l=0,u=c.length;l<u;l++)if(i=c.charCodeAt(l),i<48||i>a)return NaN;return parseInt(c,n)}}return+t};if(!d(" 0o1")||!d("0b1")||d("+0x1")){d=function(e){var t=arguments.length<1?0:e,r=this;return r instanceof d&&(v?c((function(){g.valueOf.call(r)})):o(r)!=f)?i(new b(k(t)),r,d):k(t)};for(var _,x=r("9e1e")?l(b):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),w=0;x.length>w;w++)a(b,_=x[w])&&!a(d,_)&&p(d,_,u(b,_));d.prototype=g,g.constructor=d,r("2aba")(n,f,d)}},cabd:function(e,t,r){"use strict";r.d(t,"d",(function(){return a})),r.d(t,"h",(function(){return o})),r.d(t,"c",(function(){return i})),r.d(t,"e",(function(){return s})),r.d(t,"f",(function(){return c})),r.d(t,"a",(function(){return l})),r.d(t,"b",(function(){return u})),r.d(t,"g",(function(){return p}));var n=r("b775");function a(e){return Object(n["a"])({url:"/api/v1/addcat",method:"post",data:e})}function o(){return Object(n["a"])({url:"/api/v1/getCourseCat",method:"post"})}function i(e){return Object(n["a"])({url:"/api/v1/createCourseBook",method:"post",data:e})}function s(e){return Object(n["a"])({url:"/api/v1/book/delbooks",method:"post",data:e})}function c(e){return Object(n["a"])({url:"/api/v1/book/getbooks",method:"post",data:e})}var l="https://order.meishuquanyunxiao.com/api/v1/upload",u="https://order.meishuquanyunxiao.com/admin/weChatIndex/uploadImg";function p(e){return Object(n["a"])({url:"/admin/courseDetail",method:"post",data:e})}},e659:function(e,t,r){"use strict";var n=r("6a84"),a=r.n(n);a.a},fdef:function(e,t){e.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"}}]);