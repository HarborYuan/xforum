webpackJsonp([1],{"+O20":function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("div",[a("br"),t._v(" "),a("h1",[t._v("Sign Up")]),t._v(" "),t.fail_flag?a("p",{staticStyle:{color:"#dc3545"}},[t._v("Some kinds of error")]):t._e(),t._v(" "),a("br")]),t._v(" "),a("div",{staticClass:"col-md-6 offset-md-3"},[t.show?a("b-card",{attrs:{"bg-variant":"light"}},[a("div",[a("b-form",{on:{submit:t.onSubmit,reset:t.onReset}},[a("b-form-group",{attrs:{"label-cols-sm":"5",id:"input-group-username",label:"User Name:","label-for":"input-username"}},[a("b-form-input",{attrs:{id:"input-username",type:"text",placeholder:"Enter your username",required:""},on:{input:t.onUsernameChange},model:{value:t.form.username,callback:function(e){t.$set(t.form,"username",e)},expression:"form.username"}}),t._v(" "),a("b-form-invalid-feedback",{attrs:{id:"usernameValidated",state:t.isUsernameValid}},[t._v("\n              Username Taken!\n            ")])],1),t._v(" "),a("b-form-group",{attrs:{"label-cols-sm":"5",label:"Password:","label-for":"input-password"}},[a("b-form-input",{attrs:{id:"input-password",type:"password",placeholder:"Enter your password",required:""},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1),t._v(" "),a("b-form-group",{attrs:{"label-cols-sm":"5",id:"input-group-email",label:"E-mail:","label-for":"input-email"}},[a("b-form-input",{attrs:{id:"input-email",type:"email",placeholder:"Enter your E-mail",required:""},on:{input:t.onEmailChange},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}}),t._v(" "),a("b-form-invalid-feedback",{attrs:{id:"passwordValidated",state:t.isEmailValid}},[t._v("\n              Email Registered!\n            ")])],1),t._v(" "),a("b-form-group",{staticClass:"mb-2",attrs:{"label-cols-sm":"5",id:"input-group-gender",label:"Gender:","label-for":"input-gender"}},[a("b-form-radio-group",{staticClass:"pt-2",attrs:{options:[{value:0,text:"Male"},{value:1,text:"Female"},{value:2,text:"Other"}],required:""},model:{value:t.form.gender,callback:function(e){t.$set(t.form,"gender",e)},expression:"form.gender"}})],1),t._v(" "),a("b-form-group",{staticClass:"mb-3",attrs:{"label-cols-sm":"5",id:"input-group-birthday",label:"Birthday:","label-for":"input-birthday",required:""}},[a("b-form-input",{attrs:{id:"input-birthday",type:"date",required:""},model:{value:t.form.birthday,callback:function(e){t.$set(t.form,"birthday",e)},expression:"form.birthday"}})],1),t._v(" "),a("b-form-group",{staticClass:"mb-0",attrs:{id:"submit-group"}},[a("b-button",{attrs:{type:"submit",variant:"primary"}},[t._v("Submit")]),t._v(" "),a("b-button",{attrs:{type:"reset",variant:"danger"}},[t._v("Reset")])],1)],1)],1)]):t._e()],1),t._v(" "),a("b-card",{staticClass:"mt-3",attrs:{header:"Form Data Result"}},[a("pre",{staticClass:"m-0"},[t._v(t._s(t.form))])]),t._v(" "),a("b-card",{staticClass:"mt-3",attrs:{header:"response"}},[a("pre",{staticClass:"m-0"},[t._v(t._s(t.responses))])])],1)},s=[],n={render:r,staticRenderFns:s};e.a=n},"2ml2":function(t,e,a){"use strict";var r=a("mvHQ"),s=a.n(r),n=a("mtWM"),o=a.n(n);e.a={name:"Signup",data:function(){return{form:{email:"",username:"",type:null},show_username:!1,show_email:!1}},methods:{onSubmit:function(t){t.preventDefault(),alert(s()(this.form)),o.a.post("/api/signup/",s()(this.form)).then(function(t){alert(t.data)}).catch(function(t){console.log(t.data)})},onReset:function(t){var e=this;t.preventDefault(),this.form.email="",this.form.username="",this.form.gender=null,this.form.birthday="",this.form.password="",this.show=!1,this.$nextTick(function(){e.show=!0})},onChange:function(t){t.preventDefault(),0===this.form.type?(this.show_username=!0,this.show_email=!1):(this.show_username=!1,this.show_email=!0)}}}},"94qX":function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("b-jumbotron",[a("h1",{staticClass:"display-4"},[a("span",{staticStyle:{"font-family":"'Brush Script MT'","font-size":"larger"}},[t._v("X")]),t._v("-forum")]),t._v(" "),a("p",{staticClass:"lead"},[t._v("Connect the voices around the world!")]),t._v(" "),a("b-button",{attrs:{variant:"primary",href:"#/signup"}},[t._v("Join us")]),t._v(" "),a("b-button",{attrs:{variant:"primary",href:"#/signin"}},[t._v("Sign in")])],1)],1)},s=[],n={render:r,staticRenderFns:s};e.a=n},"9M+g":function(t,e){},A93Y:function(t,e){},Asuo:function(t,e,a){"use strict";function r(t){a("DAoT")}var s=a("rOBC"),n=a("bX8z"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,"data-v-308d937d",null);e.a=l.exports},DAoT:function(t,e){},"FJN/":function(t,e,a){"use strict";function r(t){a("NwD2")}var s=a("XbJz"),n=a("94qX"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,"data-v-385f2254",null);e.a=l.exports},JjQJ:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"app"}},[a("div",[a("b-navbar",{attrs:{toggleable:"lg",type:"dark",variant:"info"}},[a("b-navbar-brand",{attrs:{href:"/"}},[t._v("X-forum")]),t._v(" "),a("b-navbar-toggle",{attrs:{target:"nav-collapse"}}),t._v(" "),a("b-collapse",{attrs:{id:"nav-collapse","is-nav":""}},[a("b-navbar-nav"),t._v(" "),t._e()],1)],1)],1),t._v(" "),a("router-view")],1)},s=[],n={render:r,staticRenderFns:s};e.a=n},KNgy:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"hello"},[a("h1",[t._v(t._s(t.msg))]),t._v(" "),a("h2",[t._v("Essential Links")]),t._v(" "),t._m(0),t._v(" "),a("h2",[t._v("Ecosystem")]),t._v(" "),t._m(1)])},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ul",[a("li",[a("a",{attrs:{href:"https://vuejs.org",target:"_blank"}},[t._v("Core Docs")])]),t._v(" "),a("li",[a("a",{attrs:{href:"https://forum.vuejs.org",target:"_blank"}},[t._v("Forum")])]),t._v(" "),a("li",[a("a",{attrs:{href:"https://chat.vuejs.org",target:"_blank"}},[t._v("Community Chat")])]),t._v(" "),a("li",[a("a",{attrs:{href:"https://twitter.com/vuejs",target:"_blank"}},[t._v("Twitter")])]),t._v(" "),a("br"),t._v(" "),a("li",[a("a",{attrs:{href:"http://vuejs-templates.github.io/webpack/",target:"_blank"}},[t._v("Docs for This Template")])])])},function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("ul",[a("li",[a("a",{attrs:{href:"http://router.vuejs.org/",target:"_blank"}},[t._v("vue-router")])]),t._v(" "),a("li",[a("a",{attrs:{href:"http://vuex.vuejs.org/",target:"_blank"}},[t._v("vuex")])]),t._v(" "),a("li",[a("a",{attrs:{href:"http://vue-loader.vuejs.org/",target:"_blank"}},[t._v("vue-loader")])]),t._v(" "),a("li",[a("a",{attrs:{href:"https://github.com/vuejs/awesome-vue",target:"_blank"}},[t._v("awesome-vue")])])])}],n={render:r,staticRenderFns:s};e.a=n},M93x:function(t,e,a){"use strict";function r(t){a("A93Y")}var s=a("xJD8"),n=a("JjQJ"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,null,null);e.a=l.exports},NHnr:function(t,e,a){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=a("7+uW"),s=a("e6fC"),n=a.n(s),o=a("qb6w"),i=(a.n(o),a("9M+g")),l=(a.n(i),a("M93x")),u=a("YaEn");r.default.use(n.a),r.default.config.productionTip=!1,new r.default({el:"#app",router:u.a,template:"<App/>",components:{App:l.a}})},NwD2:function(t,e){},OBXJ:function(t,e,a){"use strict";var r=a("mvHQ"),s=a.n(r),n=a("mtWM"),o=a.n(n),i=a("M93x");e.a={name:"Signup",data:function(){return{form:{email:"",username:"",password:"",gender:null,birthday:""},show:!0,fail_flag:!1,responses:null,isUsernameValid:"valid",isEmailValid:"valid"}},methods:{onSubmit:function(t){var e=this;t.preventDefault(),"valid"===this.isUsernameValid&&"valid"===this.isEmailValid&&(console.log("まだ？"),o.a.post(i.a.baseURL+"/api/signup/",s()(this.form)).then(function(t){console.log(t),e.responses="response","U100"===t.data&&(e.responses=t.data)}).catch(function(t){e.fail_flag=!0,console.log(t)}))},onReset:function(t){var e=this;t.preventDefault(),this.logined=!0,this.form.email="",this.form.username="",this.form.gender=null,this.form.birthday="",this.form.password="",this.show=!1,this.$nextTick(function(){e.show=!0})},onUsernameChange:function(t){var e=this;console.log(t),o.a.post(i.a.baseURL+"/api/checkusername/",s()(this.form)).then(function(t){console.log(t),"C100-1"===t.data?(console.log(t.data),e.isUsernameValid="invalid"):"C100-0"===t.data&&(e.isUsernameValid="valid")}).catch(function(t){console.log(t)})},onEmailChange:function(t){var e=this;console.log(t),o.a.post(i.a.baseURL+"/api/checkemail/",s()(this.form)).then(function(t){console.log(t),"C100-1"===t.data?e.isEmailValid="invalid":"C100-0"===t.data&&(e.isEmailValid="valid")}).catch(function(t){console.log(t)})}}}},QBed:function(t,e){},XbJz:function(t,e,a){"use strict";e.a={name:"welcome"}},YaEn:function(t,e,a){"use strict";var r=a("7+uW"),s=a("/ocq"),n=a("qSdX"),o=a("FJN/"),i=a("psOb"),l=a("Asuo"),u=a("oIpe");r.default.use(s.a),e.a=new s.a({routes:[{path:"/Hello",name:"Hello",component:n.a},{path:"/",name:"Welcome",component:o.a},{path:"/signup",name:"Signup",component:i.a},{path:"/signin",name:"Signin",component:l.a},{path:"/forget",name:"forget",component:u.a}]})},bX8z:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[t._m(0),t._v(" "),a("div",{staticClass:"col-md-6 offset-md-3"},[t.show?a("b-card",{attrs:{"bg-variant":"light"}},[a("div",[a("b-form",{on:{submit:t.onSubmit,reset:t.onReset}},[a("b-form-group",{attrs:{"label-cols-sm":"5",id:"input-group-username",label:"User Name:","label-for":"input-username"}},[a("b-form-input",{attrs:{id:"input-username",type:"text",placeholder:"Enter your username",required:""},model:{value:t.form.username,callback:function(e){t.$set(t.form,"username",e)},expression:"form.username"}})],1),t._v(" "),a("br"),t._v(" "),a("b-form-group",{attrs:{"label-cols-sm":"5",label:"Password:","label-for":"input-password"}},[a("b-form-input",{attrs:{id:"input-password",type:"password",placeholder:"Enter your password",required:""},model:{value:t.form.password,callback:function(e){t.$set(t.form,"password",e)},expression:"form.password"}})],1),t._v(" "),a("b-form-group",{staticClass:"mb-0",attrs:{id:"submit-group"}},[a("div",[a("br"),t._v(" "),a("b-button",{attrs:{type:"submit",variant:"primary"}},[t._v("Login")])],1),t._v(" "),a("div",[a("br"),t._v(" "),a("a",{attrs:{href:"#/forget"}},[t._v("Forget Password?")])]),t._v(" "),a("div")])],1)],1)]):t._e()],1),t._v(" "),a("b-card",{staticClass:"mt-3",attrs:{header:"Form Data Result"}},[a("pre",{staticClass:"m-0"},[t._v(t._s(t.form))])])],1)},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("br"),t._v(" "),a("h1",[t._v("Login in")]),t._v(" "),a("br")])}],n={render:r,staticRenderFns:s};e.a=n},"m/Wp":function(t,e){},oFch:function(t,e){},oIpe:function(t,e,a){"use strict";function r(t){a("oFch")}var s=a("2ml2"),n=a("viQp"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,"data-v-41e72b8c",null);e.a=l.exports},pMZz:function(t,e,a){"use strict";e.a={name:"hello",data:function(){return{msg:"Welcome to Your Vue.js App"}}}},psOb:function(t,e,a){"use strict";function r(t){a("m/Wp")}var s=a("OBXJ"),n=a("+O20"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,"data-v-4501bd50",null);e.a=l.exports},qSdX:function(t,e,a){"use strict";function r(t){a("QBed")}var s=a("pMZz"),n=a("KNgy"),o=a("VU/8"),i=r,l=o(s.a,n.a,!1,i,"data-v-4d10194c",null);e.a=l.exports},qb6w:function(t,e){},rOBC:function(t,e,a){"use strict";var r=a("mvHQ"),s=a.n(r),n=a("mtWM"),o=a.n(n),i=a("M93x");e.a={name:"Signup",data:function(){return{form:{username:"Bob",password:"password",response:"",base:i.a.baseURL},show:!0}},methods:{onSubmit:function(t){var e=this;t.preventDefault(),o.a.post(i.a.baseURL+"/api/signin/",s()(this.form)).then(function(t){e.form.response="Hello",e.form.res=t}).catch(function(t){e.form.response=t,console.log(t.data)})},onReset:function(t){var e=this;t.preventDefault(),this.form.email="",this.form.name="",this.form.gender=null,this.form.birthday="",this.form.password="",this.show=!1,this.$nextTick(function(){e.show=!0})}}}},viQp:function(t,e,a){"use strict";var r=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[t._m(0),t._v(" "),a("div",{staticClass:"col-md-6 offset-md-3"},[a("b-card",{attrs:{"bg-variant":"light"}},[a("div",[a("b-form",{on:{submit:t.onSubmit,reset:t.onReset,change:t.onChange}},[a("b-form-group",{staticClass:"mb-2",attrs:{"label-cols-sm":"5",id:"input-group-type",label:"What do you remember:","label-for":"input-type"}},[a("b-form-radio-group",{staticClass:"pt-2",attrs:{options:[{value:1,text:"E-mail"},{value:0,text:"Username"}]},model:{value:t.form.type,callback:function(e){t.$set(t.form,"type",e)},expression:"form.type"}})],1),t._v(" "),t.show_username?a("b-form-group",{attrs:{"label-cols-sm":"5",id:"input-group-username",label:"User Name:","label-for":"input-username"}},[a("b-form-input",{attrs:{id:"input-username",type:"text",placeholder:"Enter your username",required:""},model:{value:t.form.username,callback:function(e){t.$set(t.form,"username",e)},expression:"form.username"}})],1):t._e(),t._v(" "),t.show_email?a("b-form-group",{attrs:{"label-cols-sm":"5",id:"input-group-email",label:"E-mail:","label-for":"input-email"}},[a("b-form-input",{attrs:{id:"input-email",type:"email",placeholder:"Enter your E-mail"},model:{value:t.form.email,callback:function(e){t.$set(t.form,"email",e)},expression:"form.email"}})],1):t._e(),t._v(" "),a("b-form-group",{staticClass:"mb-0",attrs:{id:"submit-group"}},[a("b-button",{attrs:{type:"submit",variant:"primary"}},[t._v("Submit")])],1)],1)],1)])],1),t._v(" "),a("b-card",{staticClass:"mt-3",attrs:{header:"Form Data Result"}},[a("pre",{staticClass:"m-0"},[t._v(t._s(t.form))])])],1)},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("br"),t._v(" "),a("h1",[t._v("Forget")]),t._v(" "),a("br")])}],n={render:r,staticRenderFns:s};e.a=n},xJD8:function(t,e,a){"use strict";e.a={name:"app",baseURL:""}}},["NHnr"]);
//# sourceMappingURL=app.44a46e344dea7c538e45.js.map