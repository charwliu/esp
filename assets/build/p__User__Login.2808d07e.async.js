(self.webpackChunkgosh_esp=self.webpackChunkgosh_esp||[]).push([[796],{98268:function(L){L.exports={container:"container___bly96",lang:"lang___2kLqE",content:"content___17R68",top:"top___ieZWL",header:"header___3_-29",logo:"logo___St1-F",title:"title___2ZjHY",desc:"desc___1Koiu",main:"main___9TC7Z",icon:"icon___2vnhe",other:"other___NVb-Q",register:"register___NT0PJ",prefixIcon:"prefixIcon___13GU2"}},74735:function(L,T,a){"use strict";a.r(T),a.d(T,{default:function(){return se}});var ge=a(49111),$=a(19650),de=a(18106),F=a(51752),ce=a(34792),y=a(48086),h=a(11849),v=a(3182),S=a(2824),fe=a(17462),z=a(76772),D=a(94043),g=a.n(D),G=a(89366),P=a(2603),W=a(29985),R=a(36108),V=a(48107),Y=a(39464),C=a(67294),K=a(15196),N=a(5966),J=a(16434),Q=a(63434),t=a(58086),X=a(19611),w=a(29791),k=a(14146);function q(x,c){return B.apply(this,arguments)}function B(){return B=(0,v.Z)(g().mark(function x(c,f){return g().wrap(function(p){for(;;)switch(p.prev=p.next){case 0:return p.abrupt("return",(0,t.WY)("/api/login/captcha",(0,h.Z)({method:"POST",params:(0,h.Z)({},c)},f||{})));case 1:case"end":return p.stop()}},x)})),B.apply(this,arguments)}var _=a(98268),s=a.n(_),e=a(85893),I=function(c){var f=c.content;return(0,e.jsx)(z.Z,{style:{marginBottom:24},message:f,type:"error",showIcon:!0})},ee=function(){!t.m8||setTimeout(function(){var c=t.m8.location.query,f=c,M=f.redirect;t.m8.push(M||"/")},10)},ae=function(){var c=(0,C.useState)(!1),f=(0,S.Z)(c,2),M=f[0],p=f[1],te=(0,C.useState)({}),U=(0,S.Z)(te,2),H=U[0],re=U[1],ne=(0,C.useState)("account"),A=(0,S.Z)(ne,2),Z=A[0],ie=A[1],b=(0,t.tT)("@@initialState"),j=b.initialState,le=b.setInitialState,l=(0,t.YB)(),ue=function(){var d=(0,v.Z)(g().mark(function n(){var i,o;return g().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.next=2,j==null||(i=j.fetchUserInfo)===null||i===void 0?void 0:i.call(j);case 2:o=r.sent,o&&le((0,h.Z)((0,h.Z)({},j),{},{currentUser:o}));case 4:case"end":return r.stop()}},n)}));return function(){return d.apply(this,arguments)}}(),oe=function(){var d=(0,v.Z)(g().mark(function n(i){var o,m,r;return g().wrap(function(u){for(;;)switch(u.prev=u.next){case 0:return p(!0),u.prev=1,u.next=4,(0,k.x4)((0,h.Z)((0,h.Z)({},i),{},{type:Z}));case 4:if(o=u.sent,o.code!==200){u.next=12;break}return m=l.formatMessage({id:"pages.login.success",defaultMessage:"\u767B\u5F55\u6210\u529F\uFF01"}),y.default.success(m),u.next=10,ue();case 10:return ee(),u.abrupt("return");case 12:re(o),u.next=19;break;case 15:u.prev=15,u.t0=u.catch(1),r=l.formatMessage({id:"pages.login.failure",defaultMessage:"\u767B\u5F55\u5931\u8D25\uFF0C\u8BF7\u91CD\u8BD5\uFF01"}),y.default.error(r);case 19:p(!1);case 20:case"end":return u.stop()}},n,null,[[1,15]])}));return function(i){return d.apply(this,arguments)}}(),E=H.message,O=H.data;return(0,e.jsxs)("div",{className:s().container,children:[(0,e.jsx)("div",{className:s().lang,"data-lang":!0,children:t.pD&&(0,e.jsx)(t.pD,{})}),(0,e.jsxs)("div",{className:s().content,children:[(0,e.jsxs)("div",{className:s().top,children:[(0,e.jsx)("div",{className:s().header,children:(0,e.jsxs)(X.rU,{to:"/",children:[(0,e.jsx)("img",{alt:"logo",className:s().logo,src:"/logo.svg"}),(0,e.jsx)("span",{className:s().title,children:"Gosh ESP"})]})}),(0,e.jsx)("div",{className:s().desc,children:l.formatMessage({id:"pages.layouts.userLayout.title"})})]}),(0,e.jsxs)("div",{className:s().main,children:[(0,e.jsxs)(K.ZP,{initialValues:{autoLogin:!0},submitter:{searchConfig:{submitText:l.formatMessage({id:"pages.login.submit",defaultMessage:"\u767B\u5F55"})},render:function(n,i){return i.pop()},submitButtonProps:{loading:M,size:"large",style:{width:"100%"}}},onFinish:function(){var d=(0,v.Z)(g().mark(function n(i){return g().wrap(function(m){for(;;)switch(m.prev=m.next){case 0:oe(i);case 1:case"end":return m.stop()}},n)}));return function(n){return d.apply(this,arguments)}}(),children:[(0,e.jsxs)(F.Z,{activeKey:Z,onChange:ie,children:[(0,e.jsx)(F.Z.TabPane,{tab:l.formatMessage({id:"pages.login.accountLogin.tab",defaultMessage:"\u8D26\u6237\u5BC6\u7801\u767B\u5F55"})},"account"),(0,e.jsx)(F.Z.TabPane,{tab:l.formatMessage({id:"pages.login.phoneLogin.tab",defaultMessage:"\u624B\u673A\u53F7\u767B\u5F55"})},"mobile")]}),E==="error"&&O==="account"&&(0,e.jsx)(I,{content:l.formatMessage({id:"pages.login.accountLogin.errorMessage",defaultMessage:"\u8D26\u6237\u6216\u5BC6\u7801\u9519\u8BEF\uFF08admin/esp.design)"})}),Z==="account"&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(N.Z,{name:"username",fieldProps:{size:"large",prefix:(0,e.jsx)(G.Z,{className:s().prefixIcon})},placeholder:l.formatMessage({id:"pages.login.username.placeholder",defaultMessage:"\u7528\u6237\u540D: admin or user"}),rules:[{required:!0,message:(0,e.jsx)(t._H,{id:"pages.login.username.required",defaultMessage:"\u8BF7\u8F93\u5165\u7528\u6237\u540D!"})}]}),(0,e.jsx)(N.Z.Password,{name:"password",fieldProps:{size:"large",prefix:(0,e.jsx)(P.Z,{className:s().prefixIcon})},placeholder:l.formatMessage({id:"pages.login.password.placeholder",defaultMessage:"\u5BC6\u7801: esp.design"}),rules:[{required:!0,message:(0,e.jsx)(t._H,{id:"pages.login.password.required",defaultMessage:"\u8BF7\u8F93\u5165\u5BC6\u7801\uFF01"})}]})]}),E==="error"&&O==="mobile"&&(0,e.jsx)(I,{content:"\u9A8C\u8BC1\u7801\u9519\u8BEF"}),Z==="mobile"&&(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)(N.Z,{fieldProps:{size:"large",prefix:(0,e.jsx)(W.Z,{className:s().prefixIcon})},name:"mobile",placeholder:l.formatMessage({id:"pages.login.phoneNumber.placeholder",defaultMessage:"\u624B\u673A\u53F7"}),rules:[{required:!0,message:(0,e.jsx)(t._H,{id:"pages.login.phoneNumber.required",defaultMessage:"\u8BF7\u8F93\u5165\u624B\u673A\u53F7\uFF01"})},{pattern:/^1\d{10}$/,message:(0,e.jsx)(t._H,{id:"pages.login.phoneNumber.invalid",defaultMessage:"\u624B\u673A\u53F7\u683C\u5F0F\u9519\u8BEF\uFF01"})}]}),(0,e.jsx)(J.Z,{fieldProps:{size:"large",prefix:(0,e.jsx)(P.Z,{className:s().prefixIcon})},captchaProps:{size:"large"},placeholder:l.formatMessage({id:"pages.login.captcha.placeholder",defaultMessage:"\u8BF7\u8F93\u5165\u9A8C\u8BC1\u7801"}),captchaTextRender:function(n,i){return n?"".concat(i," ").concat(l.formatMessage({id:"pages.getCaptchaSecondText",defaultMessage:"\u83B7\u53D6\u9A8C\u8BC1\u7801"})):l.formatMessage({id:"pages.login.phoneLogin.getVerificationCode",defaultMessage:"\u83B7\u53D6\u9A8C\u8BC1\u7801"})},name:"captcha",rules:[{required:!0,message:(0,e.jsx)(t._H,{id:"pages.login.captcha.required",defaultMessage:"\u8BF7\u8F93\u5165\u9A8C\u8BC1\u7801\uFF01"})}],onGetCaptcha:function(){var d=(0,v.Z)(g().mark(function n(i){var o;return g().wrap(function(r){for(;;)switch(r.prev=r.next){case 0:return r.next=2,q({phone:i});case 2:if(o=r.sent,o!==!1){r.next=5;break}return r.abrupt("return");case 5:y.default.success("\u83B7\u53D6\u9A8C\u8BC1\u7801\u6210\u529F\uFF01\u9A8C\u8BC1\u7801\u4E3A\uFF1A1234");case 6:case"end":return r.stop()}},n)}));return function(n){return d.apply(this,arguments)}}()})]}),(0,e.jsxs)("div",{style:{marginBottom:24},children:[(0,e.jsx)(Q.Z,{noStyle:!0,name:"autoLogin",children:(0,e.jsx)(t._H,{id:"pages.login.rememberMe",defaultMessage:"\u81EA\u52A8\u767B\u5F55"})}),(0,e.jsx)("a",{style:{float:"right"},children:(0,e.jsx)(t._H,{id:"pages.login.forgotPassword",defaultMessage:"\u5FD8\u8BB0\u5BC6\u7801"})})]})]}),(0,e.jsxs)($.Z,{className:s().other,children:[(0,e.jsx)(t._H,{id:"pages.login.loginWith",defaultMessage:"\u5176\u4ED6\u767B\u5F55\u65B9\u5F0F"}),(0,e.jsx)(R.Z,{className:s().icon}),(0,e.jsx)(V.Z,{className:s().icon}),(0,e.jsx)(Y.Z,{className:s().icon})]})]})]}),(0,e.jsx)(w.Z,{})]})},se=ae}}]);
