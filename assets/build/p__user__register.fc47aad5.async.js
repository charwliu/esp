(self.webpackChunkgosh_esp=self.webpackChunkgosh_esp||[]).push([[803],{95526:function(g){g.exports={main:"main___2kx2N",password:"password___2JDrd",getCaptcha:"getCaptcha___1oboe",submit:"submit___34wM2",login:"login___1qBuj",success:"success___3hl98",warning:"warning___2i2r2",error:"error___3JfQo","progress-pass":"progress-pass___BM1Xu",progress:"progress___2s68u"}},83222:function(g,f,s){"use strict";s.r(f);var u=s(13062),R=s(71230),c=s(57663),B=s(71577),re=s(89032),Q=s(15746),U=s(20136),W=s(55241),Y=s(34669),te=s(32074),ie=s(34792),F=s(48086),P=s(2824),$=s(47673),d=s(4107),ae=s(43358),V=s(90290),b=s(9715),G=s(86585),y=s(67294),Z=s(58086),L=s(19611),k=s(14146),q=s(95526),E=s.n(q),e=s(85893),r=G.Z.Item,a=V.Z.Option,n=d.Z.Group,_={ok:(0,e.jsx)("div",{className:E().success,children:(0,e.jsx)("span",{children:"\u5F3A\u5EA6\uFF1A\u5F3A"})}),pass:(0,e.jsx)("div",{className:E().warning,children:(0,e.jsx)("span",{children:"\u5F3A\u5EA6\uFF1A\u4E2D"})}),poor:(0,e.jsx)("div",{className:E().error,children:(0,e.jsx)("span",{children:"\u5F3A\u5EA6\uFF1A\u592A\u77ED"})})},l={ok:"success",pass:"normal",poor:"exception"},o=function(){var p=(0,y.useState)(0),i=(0,P.Z)(p,2),h=i[0],x=i[1],M=(0,y.useState)(!1),T=(0,P.Z)(M,2),O=T[0],C=T[1],ee=(0,y.useState)("86"),S=(0,P.Z)(ee,2),I=S[0],H=S[1],X=(0,y.useState)(!1),N=(0,P.Z)(X,2),J=N[0],w=N[1],se=!1,j,z=G.Z.useForm(),_e=(0,P.Z)(z,1),A=_e[0];(0,y.useEffect)(function(){return function(){clearInterval(j)}},[j]);var le=function(){var t=59;x(t),j=window.setInterval(function(){t-=1,x(t),t===0&&clearInterval(j)},1e3)},ne=function(){var t=A.getFieldValue("password");return t&&t.length>9?"ok":t&&t.length>5?"pass":"poor"},oe=(0,Z.QT)(k.N$,{manual:!0,onSuccess:function(t,v){t.status==="ok"&&(F.default.success("\u6CE8\u518C\u6210\u529F\uFF01"),Z.m8.push({pathname:"/user/register-result",state:{account:v.email}}))}}),ue=oe.loading,ce=oe.run,de=function(t){ce(t)},ve=function(t,v){var K=Promise;return v&&v!==A.getFieldValue("password")?K.reject("\u4E24\u6B21\u8F93\u5165\u7684\u5BC6\u7801\u4E0D\u5339\u914D!"):K.resolve()},Ee=function(t,v){var K=Promise;return v?(O||C(!!v),w(!J),v.length<6?K.reject(""):(v&&se&&A.validateFields(["confirm"]),K.resolve())):(C(!!v),K.reject("\u8BF7\u8F93\u5165\u5BC6\u7801!"))},fe=function(t){H(t)},pe=function(){var t=A.getFieldValue("password"),v=ne();return t&&t.length?(0,e.jsx)("div",{className:E()["progress-".concat(v)],children:(0,e.jsx)(te.Z,{status:l[v],className:E().progress,strokeWidth:6,percent:t.length*10>100?100:t.length*10,showInfo:!1})}):null};return(0,e.jsxs)("div",{className:E().main,children:[(0,e.jsx)("h3",{children:"\u6CE8\u518C"}),(0,e.jsxs)(G.Z,{form:A,name:"UserRegister",onFinish:de,children:[(0,e.jsx)(r,{name:"mail",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u90AE\u7BB1\u5730\u5740!"},{type:"email",message:"\u90AE\u7BB1\u5730\u5740\u683C\u5F0F\u9519\u8BEF!"}],children:(0,e.jsx)(d.Z,{size:"large",placeholder:"\u90AE\u7BB1"})}),(0,e.jsx)(W.Z,{getPopupContainer:function(t){return t&&t.parentNode?t.parentNode:t},content:O&&(0,e.jsxs)("div",{style:{padding:"4px 0"},children:[_[ne()],pe(),(0,e.jsx)("div",{style:{marginTop:10},children:(0,e.jsx)("span",{children:"\u8BF7\u81F3\u5C11\u8F93\u5165 6 \u4E2A\u5B57\u7B26\u3002\u8BF7\u4E0D\u8981\u4F7F\u7528\u5BB9\u6613\u88AB\u731C\u5230\u7684\u5BC6\u7801\u3002"})})]}),overlayStyle:{width:240},placement:"right",visible:O,children:(0,e.jsx)(r,{name:"password",className:A.getFieldValue("password")&&A.getFieldValue("password").length>0&&E().password,rules:[{validator:Ee}],children:(0,e.jsx)(d.Z,{size:"large",type:"password",placeholder:"\u81F3\u5C116\u4F4D\u5BC6\u7801\uFF0C\u533A\u5206\u5927\u5C0F\u5199"})})}),(0,e.jsx)(r,{name:"confirm",rules:[{required:!0,message:"\u786E\u8BA4\u5BC6\u7801"},{validator:ve}],children:(0,e.jsx)(d.Z,{size:"large",type:"password",placeholder:"\u786E\u8BA4\u5BC6\u7801"})}),(0,e.jsxs)(n,{compact:!0,children:[(0,e.jsxs)(V.Z,{size:"large",value:I,onChange:fe,style:{width:"20%"},children:[(0,e.jsx)(a,{value:"86",children:"+86"}),(0,e.jsx)(a,{value:"87",children:"+87"})]}),(0,e.jsx)(r,{style:{width:"80%"},name:"mobile",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u624B\u673A\u53F7!"},{pattern:/^\d{11}$/,message:"\u624B\u673A\u53F7\u683C\u5F0F\u9519\u8BEF!"}],children:(0,e.jsx)(d.Z,{size:"large",placeholder:"\u624B\u673A\u53F7"})})]}),(0,e.jsxs)(R.Z,{gutter:8,children:[(0,e.jsx)(Q.Z,{span:16,children:(0,e.jsx)(r,{name:"captcha",rules:[{required:!0,message:"\u8BF7\u8F93\u5165\u9A8C\u8BC1\u7801!"}],children:(0,e.jsx)(d.Z,{size:"large",placeholder:"\u9A8C\u8BC1\u7801"})})}),(0,e.jsx)(Q.Z,{span:8,children:(0,e.jsx)(B.Z,{size:"large",disabled:!!h,className:E().getCaptcha,onClick:le,children:h?"".concat(h," s"):"\u83B7\u53D6\u9A8C\u8BC1\u7801"})})]}),(0,e.jsxs)(r,{children:[(0,e.jsx)(B.Z,{size:"large",loading:ue,className:E().submit,type:"primary",htmlType:"submit",children:(0,e.jsx)("span",{children:"\u6CE8\u518C"})}),(0,e.jsx)(L.rU,{className:E().login,to:"/user/login",children:(0,e.jsx)("span",{children:"\u4F7F\u7528\u5DF2\u6709\u8D26\u6237\u767B\u5F55"})})]})]})]})};f.default=o},15746:function(g,f,s){"use strict";var u=s(21584);f.Z=u.Z},89032:function(g,f,s){"use strict";var u=s(43673),R=s.n(u),c=s(6999)},71230:function(g,f,s){"use strict";var u=s(92820);f.Z=u.Z},13062:function(g,f,s){"use strict";var u=s(43673),R=s.n(u),c=s(6999)},19611:function(g,f,s){"use strict";s.d(f,{rU:function(){return b},OL:function(){return q}});var u=s(19630),R=s(41788),c=s(67294),B=s(97175),re=s(45697),Q=s.n(re),U=s(22122),W=s(19756),Y=s(2177),te=function(e){(0,R.Z)(r,e);function r(){for(var n,_=arguments.length,l=new Array(_),o=0;o<_;o++)l[o]=arguments[o];return n=e.call.apply(e,[this].concat(l))||this,n.history=(0,B.lX)(n.props),n}var a=r.prototype;return a.render=function(){return c.createElement(u.F0,{history:this.history,children:this.props.children})},r}(c.Component),ie=function(e){(0,R.Z)(r,e);function r(){for(var n,_=arguments.length,l=new Array(_),o=0;o<_;o++)l[o]=arguments[o];return n=e.call.apply(e,[this].concat(l))||this,n.history=(0,B.q_)(n.props),n}var a=r.prototype;return a.render=function(){return c.createElement(u.F0,{history:this.history,children:this.props.children})},r}(c.Component),F=function(r,a){return typeof r=="function"?r(a):r},P=function(r,a){return typeof r=="string"?(0,B.ob)(r,null,null,a):r},$=function(r){return r},d=c.forwardRef;typeof d=="undefined"&&(d=$);function ae(e){return!!(e.metaKey||e.altKey||e.ctrlKey||e.shiftKey)}var V=d(function(e,r){var a=e.innerRef,n=e.navigate,_=e.onClick,l=(0,W.Z)(e,["innerRef","navigate","onClick"]),o=l.target,m=(0,U.Z)({},l,{onClick:function(i){try{_&&_(i)}catch(h){throw i.preventDefault(),h}!i.defaultPrevented&&i.button===0&&(!o||o==="_self")&&!ae(i)&&(i.preventDefault(),n())}});return $!==d?m.ref=r||a:m.ref=a,c.createElement("a",m)}),b=d(function(e,r){var a=e.component,n=a===void 0?V:a,_=e.replace,l=e.to,o=e.innerRef,m=(0,W.Z)(e,["component","replace","to","innerRef"]);return c.createElement(u.s6.Consumer,null,function(p){p||(0,Y.Z)(!1);var i=p.history,h=P(F(l,p.location),p.location),x=h?i.createHref(h):"",M=(0,U.Z)({},m,{href:x,navigate:function(){var O=F(l,p.location),C=_?i.replace:i.push;C(O)}});return $!==d?M.ref=r||o:M.innerRef=o,c.createElement(n,M)})});if(!1)var G,y;var Z=function(r){return r},L=c.forwardRef;typeof L=="undefined"&&(L=Z);function k(){for(var e=arguments.length,r=new Array(e),a=0;a<e;a++)r[a]=arguments[a];return r.filter(function(n){return n}).join(" ")}var q=L(function(e,r){var a=e["aria-current"],n=a===void 0?"page":a,_=e.activeClassName,l=_===void 0?"active":_,o=e.activeStyle,m=e.className,p=e.exact,i=e.isActive,h=e.location,x=e.sensitive,M=e.strict,T=e.style,O=e.to,C=e.innerRef,ee=(0,W.Z)(e,["aria-current","activeClassName","activeStyle","className","exact","isActive","location","sensitive","strict","style","to","innerRef"]);return c.createElement(u.s6.Consumer,null,function(S){S||(0,Y.Z)(!1);var I=h||S.location,H=P(F(O,I),I),X=H.pathname,N=X&&X.replace(/([.+*?=^!:${}()[\]|/\\])/g,"\\$1"),J=N?(0,u.LX)(I.pathname,{path:N,exact:p,sensitive:x,strict:M}):null,w=!!(i?i(J,I):J),se=w?k(m,l):m,j=w?(0,U.Z)({},T,{},o):T,z=(0,U.Z)({"aria-current":w&&n||null,className:se,style:j,to:H},ee);return Z!==L?z.ref=r||C:z.innerRef=C,c.createElement(b,z)})});if(!1)var E}}]);
