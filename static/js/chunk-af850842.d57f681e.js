(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-af850842"],{"0ba0":function(t,e,n){"use strict";n.d(e,"a",(function(){return i})),n.d(e,"b",(function(){return a}));var r=n("b775");function i(t){return Object(r["a"])({url:"/admin/help/createArticle",method:"post",data:t})}function a(){return Object(r["a"])({url:"/admin/help/getArticle",method:"post"})}},"287e":function(t,e,n){},"29c6":function(t,e,n){"use strict";n.r(e);var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app-container"},[n("el-form",{ref:"ruleForm",staticClass:"demo-ruleForm",attrs:{model:t.form,rules:t.rules,"label-width":"140px","label-position":"top"}},[n("el-form-item",{attrs:{label:"主题",width:"200px",prop:"name"}},[n("el-input",{model:{value:t.form.name,callback:function(e){t.$set(t.form,"name",e)},expression:"form.name"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"子主题",width:"200px",prop:"subName"}},[n("el-input",{model:{value:t.form.subName,callback:function(e){t.$set(t.form,"subName",e)},expression:"form.subName"}})],1),t._v(" "),n("el-form-item",{attrs:{label:"分类",prop:"classifyID",width:"200px"}},[n("el-select",{attrs:{placeholder:"请选择分类"},model:{value:t.form.id,callback:function(e){t.$set(t.form,"id",e)},expression:"form.id"}},t._l(t.form.options,(function(t){return n("el-option",{key:t.id,attrs:{label:t.name,value:t.id}})})),1)],1),t._v(" "),n("el-form-item",{attrs:{label:"帮助文档详情",prop:"content"}},[n("Tinymce",{ref:"editor",attrs:{height:600},model:{value:t.form.content,callback:function(e){t.$set(t.form,"content",e)},expression:"form.content"}})],1)],1),t._v(" "),n("div",[n("div",{staticStyle:{"text-align":"right"}},[n("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.onSubmit("ruleForm")}}},[t._v("立即创建")]),t._v(" "),n("el-button",[t._v("取消")])],1)])],1)},i=[],a=(n("c5f6"),n("96cf"),n("3b8d")),o=(n("7f7f"),n("db72")),s=n("8256"),c=n("0ba0"),u=n("2f62"),l=n("52b5"),d=n("2b0e"),m={name:"createHelp",components:{Tinymce:s["a"]},computed:Object(o["a"])({},Object(u["b"])(["name"])),created:function(){this.getList()},data:function(){return{form:{name:this.$route.params.name||"",subName:this.$route.params.sub_name||"",content:this.$route.params.content||"",options:[],id:this.$route.params&&this.$route.params.classify&&this.$route.params.classify.id||""},rules:{name:[{required:!0,min:3,max:20,message:"长度在 3 到 20 个字符",trigger:"blur"}],subName:[{required:!0,min:3,max:200,message:"长度在 3 到 200 个字符",trigger:"blur"}],id:[{required:!0,message:"请选择关联分类目录",trigger:"change"}],content:[{required:!0,message:"请编辑活动内容",trigger:"blur"}]}}},methods:{getList:function(){var t=Object(a["a"])(regeneratorRuntime.mark((function t(){var e;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(l["b"])().catch((function(t){return t}));case 2:e=t.sent,2e4!==e.code&&d["default"].prototype.$message({message:e.msg,duration:1500}),this.form.options=e.data;case 5:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}(),onSubmit:function(t){var e=this;this.$route.params&&this.$route.params.id?this.$refs[t].validate(function(){var n=Object(a["a"])(regeneratorRuntime.mark((function n(r){var i,a;return regeneratorRuntime.wrap((function(n){while(1)switch(n.prev=n.next){case 0:if(!r){n.next=11;break}return i={id:e.$route.params&&Number(e.$route.params.id),name:e.form.name,sub_name:e.form.subName,content:e.form.content,classify_id:Number(e.form.id),is_edit:!0},n.next=4,Object(c["a"])(i).catch((function(t){return t}));case 4:if(a=n.sent,2e4===a.code){n.next=7;break}return n.abrupt("return",e.$message.error(a.msg));case 7:e.$refs[t].resetFields(),e.$router.go(-1),n.next=13;break;case 11:return console.log("error submit!!"),n.abrupt("return",!1);case 13:case"end":return n.stop()}}),n)})));return function(t){return n.apply(this,arguments)}}()):this.$refs[t].validate(function(){var n=Object(a["a"])(regeneratorRuntime.mark((function n(r){var i,a;return regeneratorRuntime.wrap((function(n){while(1)switch(n.prev=n.next){case 0:if(!r){n.next=11;break}return i={name:e.form.name,sub_name:e.form.subName,content:e.form.content,classify_id:Number(e.form.id),is_edit:!1},n.next=4,Object(c["a"])(i).catch((function(t){return t}));case 4:if(a=n.sent,2e4===a.code){n.next=7;break}return n.abrupt("return",e.$message.error(a.msg));case 7:e.$refs[t].resetFields(),e.$router.go(-1),n.next=13;break;case 11:return console.log("error submit!!"),n.abrupt("return",!1);case 13:case"end":return n.stop()}}),n)})));return function(t){return n.apply(this,arguments)}}())}}},f=m,h=n("2877"),p=Object(h["a"])(f,r,i,!1,null,"721b4f85",null);e["default"]=p.exports},"456d":function(t,e,n){var r=n("4bf8"),i=n("0d58");n("5eda")("keys",(function(){return function(t){return i(r(t))}}))},"52b5":function(t,e,n){"use strict";n.d(e,"a",(function(){return i})),n.d(e,"b",(function(){return a})),n.d(e,"c",(function(){return o}));var r=n("b775");function i(t){return Object(r["a"])({url:"/classify/createClassify",method:"post",data:t})}function a(t){return Object(r["a"])({url:"/classify/getClassify",method:"post",data:t})}function o(t){return Object(r["a"])({url:"/classify/upClassify",method:"post",data:t})}},"5eda":function(t,e,n){var r=n("5ca1"),i=n("8378"),a=n("79e5");t.exports=function(t,e){var n=(i.Object||{})[t]||Object[t],o={};o[t]=e(n),r(r.S+r.F*a((function(){n(1)})),"Object",o)}},7654:function(t,e,n){"use strict";var r=n("287e"),i=n.n(r);i.a},8256:function(t,e,n){"use strict";var r=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"tinymce-container",class:{fullscreen:t.fullscreen},style:{width:t.containerWidth}},[n("textarea",{staticClass:"tinymce-textarea",attrs:{id:t.tinymceId}}),t._v(" "),n("div",{staticClass:"editor-custom-btn-container"},[n("editorImage",{staticClass:"editor-upload-btn",attrs:{color:"#1890ff"},on:{successCBK:t.imageSuccessCBK}})],1)])},i=[],a=(n("ac6a"),n("c5f6"),function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"upload-container"},[n("el-button",{style:{background:t.color,borderColor:t.color},attrs:{icon:"el-icon-upload",size:"mini",type:"primary"},on:{click:function(e){t.dialogVisible=!0}}},[t._v("\n    本地图片上传\n  ")]),t._v(" "),n("el-dialog",{attrs:{visible:t.dialogVisible},on:{"update:visible":function(e){t.dialogVisible=e}}},[n("el-upload",{staticClass:"editor-slide-upload",attrs:{name:"upload[]",multiple:!0,"file-list":t.fileList,"show-file-list":!0,"on-remove":t.handleRemove,"on-success":t.handleSuccess,"before-upload":t.beforeUpload,action:t.upImgURL,"list-type":"picture-card"}},[n("el-button",{attrs:{size:"small",type:"primary"}},[t._v("\n        点击 上传\n      ")])],1),t._v(" "),n("el-button",{on:{click:function(e){t.dialogVisible=!1}}},[t._v("\n      取消\n    ")]),t._v(" "),n("el-button",{attrs:{type:"primary"},on:{click:t.handleSubmit}},[t._v("\n      确认\n    ")])],1)],1)}),o=[],s=(n("456d"),n("cabd")),c={name:"EditorSlideUpload",props:{color:{type:String,default:"#1890ff"}},data:function(){return{upImgURL:s["a"],dialogVisible:!1,listObj:{},fileList:[]}},methods:{checkAllSuccess:function(){var t=this;return Object.keys(this.listObj).every((function(e){return t.listObj[e].hasSuccess}))},handleSubmit:function(){var t=this,e=Object.keys(this.listObj).map((function(e){return t.listObj[e]}));this.checkAllSuccess()?(this.$emit("successCBK",e),this.listObj={},this.fileList=[],this.dialogVisible=!1):this.$message("请等待所有图像成功上载。如果有网络问题，请刷新页面并重新上传！")},handleSuccess:function(t,e,n){console.log(t,"file------");for(var r=e.uid,i=Object.keys(this.listObj),a=0,o=i.length;a<o;a++)if(this.listObj[i[a]].uid===r)return this.listObj[i[a]].url=t.files.file,void(this.listObj[i[a]].hasSuccess=!0)},handleRemove:function(t){for(var e=t.uid,n=Object.keys(this.listObj),r=0,i=n.length;r<i;r++)if(this.listObj[n[r]].uid===e)return void delete this.listObj[n[r]]},beforeUpload:function(t){var e=this,n=this,r=window.URL||window.webkitURL,i=t.uid;return this.listObj[i]={},new Promise((function(a,o){var s=new Image;s.src=r.createObjectURL(t),s.onload=function(){n.listObj[i]={hasSuccess:!1,uid:t.uid,width:this.width,height:this.height}},console.log(e.listObj,"---------log------"),a(!0)}))}}},u=c,l=(n("7654"),n("2877")),d=Object(l["a"])(u,a,o,!1,null,"0b4d82ca",null),m=d.exports,f=["advlist anchor autolink autosave code codesample    directionality emoticons fullscreen hr image imagetools insertdatetime link lists media nonbreaking noneditable pagebreak paste preview print save searchreplace  tabfocus table template  textpattern visualblocks visualchars wordcount toc fontselect lineheight "],h=f,p=["searchreplace bold italic underline strikethrough alignleft aligncenter alignright outdent indent  blockquote undo redo removeformat subscript superscript code codesample","hr lineheight fontselect toc fontsizeselect formatselect removeformat bullist numlist link image charmap preview anchor pagebreak insertdatetime media table emoticons forecolor backcolor fullscreen"],b=p,g=n("2d63"),v=[];function y(){return window.tinymce}var w=function(t,e){var n=document.getElementById(t),r=e||function(){};if(!n){var i=document.createElement("script");i.src=t,i.id=t,document.body.appendChild(i),v.push(r);var a="onload"in i?o:s;a(i)}function o(e){e.onload=function(){this.onerror=this.onload=null;var t,n=Object(g["a"])(v);try{for(n.s();!(t=n.n()).done;){var r=t.value;r(null,e)}}catch(i){n.e(i)}finally{n.f()}v=null},e.onerror=function(){this.onerror=this.onload=null,r(new Error("Failed to load "+t),e)}}function s(t){t.onreadystatechange=function(){if("complete"===this.readyState||"loaded"===this.readyState){this.onreadystatechange=null;var e,n=Object(g["a"])(v);try{for(n.s();!(e=n.n()).done;){var r=e.value;r(null,t)}}catch(i){n.e(i)}finally{n.f()}v=null}}}n&&r&&(y()?r(null,n):v.push(r))},_=w,k="https://order.meishuquanyunxiao.com/myPackge/tinymce.min.js",x={name:"Tinymce",components:{editorImage:m},props:{id:{type:String,default:function(){return"vue-tinymce-"+ +new Date+(1e3*Math.random()).toFixed(0)}},value:{type:String,default:""},toolbar:{type:Array,required:!1,default:function(){return[]}},menubar:{type:String,default:"file edit insert view format table"},height:{type:[Number,String],required:!1,default:360},width:{type:[Number,String],required:!1,default:"auto"}},data:function(){return{hasChange:!1,hasInit:!1,tinymceId:this.id,fullscreen:!1,languageTypeList:{zh:"zh_CN"}}},computed:{containerWidth:function(){var t=this.width;return/^[\d]+(\.[\d]+)?$/.test(t)?"".concat(t,"px"):t}},watch:{value:function(t){var e=this;!this.hasChange&&this.hasInit&&this.$nextTick((function(){return window.tinymce.get(e.tinymceId).setContent(t||"")}))}},mounted:function(){this.init()},activated:function(){window.tinymce&&this.initTinymce()},deactivated:function(){this.destroyTinymce()},destroyed:function(){this.destroyTinymce()},methods:{init:function(){var t=this;_(k,(function(e){e?t.$message.error(e.message):t.initTinymce()}))},initTinymce:function(){var t=this,e=this;window.tinymce.init({selector:"#".concat(this.tinymceId),language_url:"https://order.meishuquanyunxiao.com/myPackge/langs/zh_CN.js",language:this.languageTypeList["zh"],height:this.height,branding:!1,body_class:"panel-body ",object_resizing:!0,toolbar:this.toolbar.length>0?this.toolbar:b,menubar:this.menubar,plugins:h,end_container_on_empty_block:!0,powerpaste_word_import:"clean",textpattern_patterns:[{start:"*",end:"*",format:"italic"},{start:"**",end:"**",format:"bold"},{start:"#",format:"h1"},{start:"##",format:"h2"},{start:"###",format:"h3"},{start:"####",format:"h4"},{start:"#####",format:"h5"},{start:"######",format:"h6"},{start:"1. ",cmd:"InsertOrderedList"},{start:"* ",cmd:"InsertUnorderedList"},{start:"- ",cmd:"InsertUnorderedList"}],fontsize_formats:"8px 10px 12px 14px 18px 24px 36px",font_formats:"微软雅黑='微软雅黑';宋体='宋体';黑体='黑体';仿宋='仿宋';楷体='楷体';隶书='隶书';幼圆='幼圆';Andale Mono=andale mono,times;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;Comic Sans MS=comic sans ms,sans-serif;Courier New=courier new,courier;Georgia=georgia,palatino;Helvetica=helvetica;Impact=impact,chicago;Symbol=symbol;Tahoma=tahoma,arial,helvetica,sans-serif;Terminal=terminal,monaco;Times New Roman=times new roman,times;Trebuchet MS=trebuchet ms,geneva;Verdana=verdana,geneva;Webdings=webdings;Wingdings=wingdings",code_dialog_height:450,code_dialog_width:1e3,advlist_bullet_styles:"default,circle,disc,square",advlist_number_styles:"default,lower-alpha,lower-greek,lower-roman,upper-alpha,upper-roman",imagetools_cors_hosts:["www.tinymce.com","codepen.io","127.0.0.1:8080"],default_link_target:"_blank",link_title:!1,nonbreaking_force_tab:!0,init_instance_callback:function(n){e.value&&n.setContent(e.value),e.hasInit=!0,n.on("NodeChange Change KeyUp SetContent",(function(){t.hasChange=!0,t.$emit("input",n.getContent())}))},setup:function(t){t.on("FullscreenStateChanged",(function(t){e.fullscreen=t.state}))},convert_urls:!1})},destroyTinymce:function(){var t=window.tinymce.get(this.tinymceId);this.fullscreen&&t.execCommand("mceFullScreen"),t&&t.destroy()},setContent:function(t){window.tinymce.get(this.tinymceId).setContent(t)},getContent:function(){window.tinymce.get(this.tinymceId).getContent()},imageSuccessCBK:function(t){var e=this;t.forEach((function(t){window.tinymce.get(e.tinymceId).insertContent('<img class="wscnph" src="'.concat(t.url,'" >'))}))}}},O=x,j=(n("8389"),Object(l["a"])(O,r,i,!1,null,"d0342d0a",null));e["a"]=j.exports},8389:function(t,e,n){"use strict";var r=n("da5e"),i=n.n(r);i.a},aa77:function(t,e,n){var r=n("5ca1"),i=n("be13"),a=n("79e5"),o=n("fdef"),s="["+o+"]",c="​",u=RegExp("^"+s+s+"*"),l=RegExp(s+s+"*$"),d=function(t,e,n){var i={},s=a((function(){return!!o[t]()||c[t]()!=c})),u=i[t]=s?e(m):o[t];n&&(i[n]=u),r(r.P+r.F*s,"String",i)},m=d.trim=function(t,e){return t=String(i(t)),1&e&&(t=t.replace(u,"")),2&e&&(t=t.replace(l,"")),t};t.exports=d},c5f6:function(t,e,n){"use strict";var r=n("7726"),i=n("69a8"),a=n("2d95"),o=n("5dbc"),s=n("6a99"),c=n("79e5"),u=n("9093").f,l=n("11e9").f,d=n("86cc").f,m=n("aa77").trim,f="Number",h=r[f],p=h,b=h.prototype,g=a(n("2aeb")(b))==f,v="trim"in String.prototype,y=function(t){var e=s(t,!1);if("string"==typeof e&&e.length>2){e=v?e.trim():m(e,3);var n,r,i,a=e.charCodeAt(0);if(43===a||45===a){if(n=e.charCodeAt(2),88===n||120===n)return NaN}else if(48===a){switch(e.charCodeAt(1)){case 66:case 98:r=2,i=49;break;case 79:case 111:r=8,i=55;break;default:return+e}for(var o,c=e.slice(2),u=0,l=c.length;u<l;u++)if(o=c.charCodeAt(u),o<48||o>i)return NaN;return parseInt(c,r)}}return+e};if(!h(" 0o1")||!h("0b1")||h("+0x1")){h=function(t){var e=arguments.length<1?0:t,n=this;return n instanceof h&&(g?c((function(){b.valueOf.call(n)})):a(n)!=f)?o(new p(y(e)),n,h):y(e)};for(var w,_=n("9e1e")?u(p):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),k=0;_.length>k;k++)i(p,w=_[k])&&!i(h,w)&&d(h,w,l(p,w));h.prototype=b,b.constructor=h,n("2aba")(r,f,h)}},cabd:function(t,e,n){"use strict";n.d(e,"d",(function(){return i})),n.d(e,"h",(function(){return a})),n.d(e,"c",(function(){return o})),n.d(e,"e",(function(){return s})),n.d(e,"f",(function(){return c})),n.d(e,"a",(function(){return u})),n.d(e,"b",(function(){return l})),n.d(e,"g",(function(){return d}));var r=n("b775");function i(t){return Object(r["a"])({url:"/api/v1/addcat",method:"post",data:t})}function a(){return Object(r["a"])({url:"/api/v1/getCourseCat",method:"post"})}function o(t){return Object(r["a"])({url:"/api/v1/createCourseBook",method:"post",data:t})}function s(t){return Object(r["a"])({url:"/api/v1/book/delbooks",method:"post",data:t})}function c(t){return Object(r["a"])({url:"/api/v1/book/getbooks",method:"post",data:t})}var u="https://order.meishuquanyunxiao.com/api/v1/upload",l="https://order.meishuquanyunxiao.com/admin/weChatIndex/uploadImg";function d(t){return Object(r["a"])({url:"/admin/courseDetail",method:"post",data:t})}},da5e:function(t,e,n){},fdef:function(t,e){t.exports="\t\n\v\f\r   ᠎             　\u2028\u2029\ufeff"}}]);