(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0aef37"],{"0bf5":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"container"},[a("h3",[t._v("All Holidays")]),t.message?a("div",{staticClass:"alert alert-success"},[t._v("\n    "+t._s(this.message))]):t._e(),a("div",{staticClass:"container"},[a("table",{staticClass:"table"},[t._m(0),a("tbody",t._l(t.holidays,function(e){return a("tr",{key:e.id},[a("td",[t._v(t._s(e.name))]),a("td",[t._v(t._s(e.description))]),a("td",[t._v(t._s(e.date))]),a("td",[a("button",{staticClass:"btn btn-warning",on:{click:function(a){return t.updateHoliday(e.id)}}},[t._v("\n              Update\n            ")])]),a("td",[a("button",{staticClass:"btn btn-danger",on:{click:function(a){return t.deleteHoliday(e.id)}}},[t._v("\n              Delete\n            ")])])])}),0)]),a("div",{staticClass:"row"},[a("button",{staticClass:"btn btn-success",on:{click:function(e){return t.addHoliday()}}},[t._v("Add")])])])])},s=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("thead",[a("tr",[a("th",[t._v("Name")]),a("th",[t._v("Description")]),a("th",[t._v("Date")]),a("th",[t._v("Update")]),a("th",[t._v("Delete")])])])}],i=a("a6a3"),d={name:"Holidays",data:function(){return{holidays:[],message:""}},methods:{refreshHolidays:function(){var t=this;i["a"].retrieveAllHolidays().then(function(e){t.holidays=e.data})},addHoliday:function(){this.$router.push("/holiday/-1")},updateHoliday:function(t){this.$router.push("/holiday/".concat(t))},deleteHoliday:function(t){var e=this;i["a"].deleteHoliday(t).then(function(){e.refreshHolidays()})}},created:function(){this.refreshHolidays()}},l=d,o=a("2877"),c=Object(o["a"])(l,n,s,!1,null,null,null);e["default"]=c.exports}}]);
//# sourceMappingURL=chunk-2d0aef37.d10eef5e.js.map