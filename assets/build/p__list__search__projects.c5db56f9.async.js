(self.webpackChunkgosh_esp=self.webpackChunkgosh_esp||[]).push([[33],{99165:function(j,Z,e){"use strict";e.d(Z,{Z:function(){return A}});var u=e(67294),T={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M832 64H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496v688c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V96c0-17.7-14.3-32-32-32zM704 192H192c-17.7 0-32 14.3-32 32v530.7c0 8.5 3.4 16.6 9.4 22.6l173.3 173.3c2.2 2.2 4.7 4 7.4 5.5v1.9h4.2c3.5 1.3 7.2 2 11 2H704c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32zM350 856.2L263.9 770H350v86.2zM664 888H414V746c0-22.1-17.9-40-40-40H232V264h432v624z"}}]},name:"copy",theme:"outlined"},L=T,E=e(27029),v=function(a,n){return u.createElement(E.Z,Object.assign({},a,{ref:n,icon:L}))};v.displayName="CopyOutlined";var A=u.forwardRef(v)},58491:function(j,Z,e){"use strict";e.d(Z,{Z:function(){return A}});var u=e(67294),T={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M890.5 755.3L537.9 269.2c-12.8-17.6-39-17.6-51.7 0L133.5 755.3A8 8 0 00140 768h75c5.1 0 9.9-2.5 12.9-6.6L512 369.8l284.1 391.6c3 4.1 7.8 6.6 12.9 6.6h75c6.5 0 10.3-7.4 6.5-12.7z"}}]},name:"up",theme:"outlined"},L=T,E=e(27029),v=function(a,n){return u.createElement(E.Z,Object.assign({},a,{ref:n,icon:L}))};v.displayName="UpOutlined";var A=u.forwardRef(v)},24444:function(j){j.exports={avatarList:"avatarList___1Twgv",avatarItem:"avatarItem___16EyN",avatarItemLarge:"avatarItemLarge___5VUZ6",avatarItemSmall:"avatarItemSmall___rCe9R",avatarItemMini:"avatarItemMini___2fmPX"}},10991:function(j){j.exports={standardFormRow:"standardFormRow___rVZMU",label:"label___2UOXv",content:"content___pJkbN",standardFormRowLast:"standardFormRowLast___1apgA",standardFormRowBlock:"standardFormRowBlock___eVu8k",standardFormRowGrid:"standardFormRowGrid___3aiHp"}},39843:function(j){j.exports={tagSelect:"tagSelect___1aaMH",expanded:"expanded___3hv8W",trigger:"trigger___3t1ed",anticon:"anticon___12aY1",hasExpandTag:"hasExpandTag___1WI1K"}},45168:function(j){j.exports={coverCardList:"coverCardList___2LrlR",card:"card___1WgqT",cardItemContent:"cardItemContent___Un4wM",avatarList:"avatarList___2kgtw",cardList:"cardList___2OFVD"}},85637:function(j,Z,e){"use strict";e.r(Z),e.d(Z,{default:function(){return Me}});var u=e(13062),T=e(71230),L=e(89032),E=e(15746),v=e(11849),A=e(54421),r=e(38272),a=e(72012),n=e(39144),s=e(402),o=e(59289),m=e(9715),d=e(86585),c=e(43358),f=e(90290),P=e(14146),C=e(30381),I=e.n(C),M=e(58086),X=e(93224),ne=e(22385),Q=e(69713),q=e(94233),U=e(51890),D=e(32059),W=e(67294),Ce=e(94184),_=e.n(Ce),je=e(24444),J=e.n(je),t=e(85893),re=function(l){var i;return _()(J().avatarItem,(i={},(0,D.Z)(i,J().avatarItemLarge,l==="large"),(0,D.Z)(i,J().avatarItemSmall,l==="small"),(0,D.Z)(i,J().avatarItemMini,l==="mini"),i))},Ee=function(l){var i=l.src,h=l.size,y=l.tips,g=l.onClick,R=g===void 0?function(){}:g,S=re(h);return(0,t.jsx)("li",{className:S,onClick:R,children:y?(0,t.jsx)(Q.Z,{title:y,children:(0,t.jsx)(U.C,{src:i,size:h,style:{cursor:"pointer"}})}):(0,t.jsx)(U.C,{src:i,size:h})})},le=function(l){var i=l.children,h=l.size,y=l.maxLength,g=y===void 0?5:y,R=l.excessItemsStyle,S=(0,X.Z)(l,["children","size","maxLength","excessItemsStyle"]),O=W.Children.count(i),x=g>=O?O:g,F=W.Children.toArray(i),$=F.slice(0,x).map(function(k){return W.cloneElement(k,{size:h})});if(x<O){var w=re(h);$.push((0,t.jsx)("li",{className:w,children:(0,t.jsx)(U.C,{size:h,style:R,children:"+".concat(O-g)})},"exceed"))}return(0,t.jsx)("div",(0,v.Z)((0,v.Z)({},S),{},{className:J().avatarList,children:(0,t.jsxs)("ul",{children:[" ",$," "]})}))};le.Item=Ee;var se=le,pe=e(10991),H=e.n(pe),Oe=function(l){var i,h=l.title,y=l.children,g=l.last,R=l.block,S=l.grid,O=(0,X.Z)(l,["title","children","last","block","grid"]),x=_()(H().standardFormRow,(i={},(0,D.Z)(i,H().standardFormRowBlock,R),(0,D.Z)(i,H().standardFormRowLast,g),(0,D.Z)(i,H().standardFormRowGrid,S),i));return(0,t.jsxs)("div",(0,v.Z)((0,v.Z)({className:x},O),{},{children:[h&&(0,t.jsx)("div",{className:H().label,children:(0,t.jsx)("span",{children:h})}),(0,t.jsx)("div",{className:H().content,children:y})]}))},oe=Oe,Ze=e(86582),ie=e(2824),Ge=e(71153),Te=e(60331),Se=e(58491),Le=e(57254),Ie=e(59819),Be=e(74081),Ne=e(39843),Y=e.n(Ne),de=Te.Z.CheckableTag,ce=function(l){var i=l.children,h=l.checked,y=l.onChange,g=l.value;return(0,t.jsx)(de,{checked:!!h,onChange:function(S){return y&&y(g,S)},children:i},g)};ce.isTagSelectOption=!0;var ue=function(l){var i,h=l.children,y=l.hideCheckAll,g=y===void 0?!1:y,R=l.className,S=l.style,O=l.expandable,x=l.actionsText,F=x===void 0?{}:x,$=(0,Ie.Z)(),w=(0,ie.Z)($,2),k=w[0],Ue=w[1].toggle,De=(0,Be.Z)(l),ve=(0,ie.Z)(De,2),G=ve[0],me=ve[1],he=function(N){return N&&N.type&&(N.type.isTagSelectOption||N.type.displayName==="TagSelectOption")},fe=function(){var N=W.Children.toArray(h),V=N.filter(function(z){return he(z)}).map(function(z){return z.props.value});return V||[]},ze=function(N){var V=[];N&&(V=fe()),me(V)},Ke=function(N,V){var z=(0,Ze.Z)(G||[]),ae=z.indexOf(N);V&&ae===-1?z.push(N):!V&&ae>-1&&z.splice(ae,1),me(z)},Ve=fe().length===(G==null?void 0:G.length),ge=F.expandText,We=ge===void 0?"\u5C55\u5F00":ge,xe=F.collapseText,He=xe===void 0?"\u6536\u8D77":xe,ye=F.selectAllText,be=ye===void 0?"\u5168\u90E8":ye,$e=_()(Y().tagSelect,R,(i={},(0,D.Z)(i,Y().hasExpandTag,O),(0,D.Z)(i,Y().expanded,k),i));return(0,t.jsxs)("div",{className:$e,style:S,children:[g?null:(0,t.jsx)(de,{checked:Ve,onChange:ze,children:be},"tag-select-__all__"),h&&W.Children.map(h,function(B){return he(B)?W.cloneElement(B,{key:"tag-select-".concat(B.props.value),value:B.props.value,checked:G&&G.indexOf(B.props.value)>-1,onChange:Ke}):B}),O&&(0,t.jsx)("a",{className:Y().trigger,onClick:function(){Ue()},children:k?(0,t.jsxs)(t.Fragment,{children:[He," ",(0,t.jsx)(Se.Z,{})]}):(0,t.jsxs)(t.Fragment,{children:[We,(0,t.jsx)(Le.Z,{})]})})]})};ue.Option=ce;var p=ue,Ae=e(45168),b=e.n(Ae),ee=f.Z.Option,te=d.Z.Item,Re=o.Z.Paragraph,Fe=function(l,i){return"".concat(l,"-").concat(i)},Pe=function(){var l=(0,M.QT)(function(O){return console.log("form data",O),(0,P.Gv)({count:8})}),i=l.data,h=l.loading,y=l.run,g=(i==null?void 0:i.list)||[],R=g&&(0,t.jsx)(r.ZP,{rowKey:"id",loading:h,grid:{gutter:16,xs:1,sm:2,md:3,lg:3,xl:4,xxl:4},dataSource:g,renderItem:function(x){return(0,t.jsx)(r.ZP.Item,{children:(0,t.jsxs)(n.Z,{className:b().card,hoverable:!0,cover:(0,t.jsx)("img",{alt:x.title,src:x.cover}),children:[(0,t.jsx)(n.Z.Meta,{title:(0,t.jsx)("a",{children:x.title}),description:(0,t.jsx)(Re,{className:b().item,ellipsis:{rows:2},children:x.subDescription})}),(0,t.jsxs)("div",{className:b().cardItemContent,children:[(0,t.jsx)("span",{children:I()(x.updatedAt).fromNow()}),(0,t.jsx)("div",{className:b().avatarList,children:(0,t.jsx)(se,{size:"small",children:x.members.map(function(F,$){return(0,t.jsx)(se.Item,{src:F.avatar,tips:F.name},Fe(x.id,$))})})})]})]})})}}),S={wrapperCol:{xs:{span:24},sm:{span:16}}};return(0,t.jsxs)("div",{className:b().coverCardList,children:[(0,t.jsx)(n.Z,{bordered:!1,children:(0,t.jsxs)(d.Z,{layout:"inline",onValuesChange:function(x,F){y(F)},children:[(0,t.jsx)(oe,{title:"\u6240\u5C5E\u7C7B\u76EE",block:!0,style:{paddingBottom:11},children:(0,t.jsx)(te,{name:"category",children:(0,t.jsxs)(p,{expandable:!0,children:[(0,t.jsx)(p.Option,{value:"cat1",children:"\u7C7B\u76EE\u4E00"}),(0,t.jsx)(p.Option,{value:"cat2",children:"\u7C7B\u76EE\u4E8C"}),(0,t.jsx)(p.Option,{value:"cat3",children:"\u7C7B\u76EE\u4E09"}),(0,t.jsx)(p.Option,{value:"cat4",children:"\u7C7B\u76EE\u56DB"}),(0,t.jsx)(p.Option,{value:"cat5",children:"\u7C7B\u76EE\u4E94"}),(0,t.jsx)(p.Option,{value:"cat6",children:"\u7C7B\u76EE\u516D"}),(0,t.jsx)(p.Option,{value:"cat7",children:"\u7C7B\u76EE\u4E03"}),(0,t.jsx)(p.Option,{value:"cat8",children:"\u7C7B\u76EE\u516B"}),(0,t.jsx)(p.Option,{value:"cat9",children:"\u7C7B\u76EE\u4E5D"}),(0,t.jsx)(p.Option,{value:"cat10",children:"\u7C7B\u76EE\u5341"}),(0,t.jsx)(p.Option,{value:"cat11",children:"\u7C7B\u76EE\u5341\u4E00"}),(0,t.jsx)(p.Option,{value:"cat12",children:"\u7C7B\u76EE\u5341\u4E8C"})]})})}),(0,t.jsx)(oe,{title:"\u5176\u5B83\u9009\u9879",grid:!0,last:!0,children:(0,t.jsxs)(T.Z,{gutter:16,children:[(0,t.jsx)(E.Z,{lg:8,md:10,sm:10,xs:24,children:(0,t.jsx)(te,(0,v.Z)((0,v.Z)({},S),{},{label:"\u4F5C\u8005",name:"author",children:(0,t.jsx)(f.Z,{placeholder:"\u4E0D\u9650",style:{maxWidth:200,width:"100%"},children:(0,t.jsx)(ee,{value:"lisa",children:"\u738B\u662D\u541B"})})}))}),(0,t.jsx)(E.Z,{lg:8,md:10,sm:10,xs:24,children:(0,t.jsx)(te,(0,v.Z)((0,v.Z)({},S),{},{label:"\u597D\u8BC4\u5EA6",name:"rate",children:(0,t.jsxs)(f.Z,{placeholder:"\u4E0D\u9650",style:{maxWidth:200,width:"100%"},children:[(0,t.jsx)(ee,{value:"good",children:"\u4F18\u79C0"}),(0,t.jsx)(ee,{value:"normal",children:"\u666E\u901A"})]})}))})]})})]})}),(0,t.jsx)("div",{className:b().cardList,children:R})]})},Me=Pe},59819:function(j,Z,e){"use strict";e.d(Z,{Z:function(){return A}});var u=e(67294),T=function(r,a){var n=typeof Symbol=="function"&&r[Symbol.iterator];if(!n)return r;var s=n.call(r),o,m=[],d;try{for(;(a===void 0||a-- >0)&&!(o=s.next()).done;)m.push(o.value)}catch(c){d={error:c}}finally{try{o&&!o.done&&(n=s.return)&&n.call(s)}finally{if(d)throw d.error}}return m};function L(r,a){r===void 0&&(r=!1);var n=T((0,u.useState)(r),2),s=n[0],o=n[1],m=(0,u.useMemo)(function(){var d=a===void 0?!r:a,c=function(I){if(I!==void 0){o(I);return}o(function(M){return M===r?d:r})},f=function(){return o(r)},P=function(){return o(d)};return{toggle:c,setLeft:f,setRight:P}},[r,a]);return[s,m]}var E=L,v=function(r,a){var n=typeof Symbol=="function"&&r[Symbol.iterator];if(!n)return r;var s=n.call(r),o,m=[],d;try{for(;(a===void 0||a-- >0)&&!(o=s.next()).done;)m.push(o.value)}catch(c){d={error:c}}finally{try{o&&!o.done&&(n=s.return)&&n.call(s)}finally{if(d)throw d.error}}return m};function A(r){r===void 0&&(r=!1);var a=v(E(r),2),n=a[0],s=a[1].toggle,o=(0,u.useMemo)(function(){var m=function(){return s(!0)},d=function(){return s(!1)};return{toggle:s,setTrue:m,setFalse:d}},[s]);return[n,o]}},74081:function(j,Z,e){"use strict";e.d(Z,{Z:function(){return r}});var u=e(67294),T=function(n,s){var o=(0,u.useRef)(!1);(0,u.useEffect)(function(){if(!o.current)o.current=!0;else return n()},s)},L=T,E=function(a,n){var s=typeof Symbol=="function"&&a[Symbol.iterator];if(!s)return a;var o=s.call(a),m,d=[],c;try{for(;(n===void 0||n-- >0)&&!(m=o.next()).done;)d.push(m.value)}catch(f){c={error:f}}finally{try{m&&!m.done&&(s=o.return)&&s.call(o)}finally{if(c)throw c.error}}return d},v=function(){for(var a=[],n=0;n<arguments.length;n++)a=a.concat(E(arguments[n]));return a};function A(a,n){a===void 0&&(a={}),n===void 0&&(n={});var s=n.defaultValue,o=n.defaultValuePropName,m=o===void 0?"defaultValue":o,d=n.valuePropName,c=d===void 0?"value":d,f=n.trigger,P=f===void 0?"onChange":f,C=a[c],I=E((0,u.useState)(function(){return c in a?C:m in a?a[m]:s}),2),M=I[0],X=I[1];L(function(){c in a&&X(C)},[C,c]);var ne=(0,u.useCallback)(function(Q){for(var q=[],U=1;U<arguments.length;U++)q[U-1]=arguments[U];c in a||X(Q),a[P]&&a[P].apply(a,v([Q],q))},[a,c,P]);return[c in a?C:M,ne]}var r=A},34952:function(j,Z,e){"use strict";var u=e(22122),T=e(67294),L=e(15105),E=function(r,a){var n={};for(var s in r)Object.prototype.hasOwnProperty.call(r,s)&&a.indexOf(s)<0&&(n[s]=r[s]);if(r!=null&&typeof Object.getOwnPropertySymbols=="function")for(var o=0,s=Object.getOwnPropertySymbols(r);o<s.length;o++)a.indexOf(s[o])<0&&Object.prototype.propertyIsEnumerable.call(r,s[o])&&(n[s[o]]=r[s[o]]);return n},v={border:0,background:"transparent",padding:0,lineHeight:"inherit",display:"inline-block"},A=T.forwardRef(function(r,a){var n=function(C){var I=C.keyCode;I===L.Z.ENTER&&C.preventDefault()},s=function(C){var I=C.keyCode,M=r.onClick;I===L.Z.ENTER&&M&&M()},o=r.style,m=r.noStyle,d=r.disabled,c=E(r,["style","noStyle","disabled"]),f={};return m||(f=(0,u.Z)({},v)),d&&(f.pointerEvents="none"),f=(0,u.Z)((0,u.Z)({},f),o),T.createElement("div",(0,u.Z)({role:"button",tabIndex:0,ref:a},c,{onKeyDown:n,onKeyUp:s,style:f}))});Z.Z=A}}]);
