/*
* @Author: layui-2
* @Date:   2018-08-31 11:40:42
* @Last Modified by:   layui-2
* @Last Modified time: 2018-09-04 14:44:38
*/
layui.define([], function (exports) {
	"use strict";
	var $ = layui.jquery

		//外部接口
		, layTags = {
			config: {}

			//设置全局项
			, set: function (options) {
				var that = this;
				that.config = $.extend({}, that.config, options);
				return that;
			}

			// 事件监听
			, on: function (events, callback) {
				return layui.onevent.call(this, MOD_NAME, events, callback)
			}
		}

		//操作当前实例
		, thisLayTags = function () {
			var that = this
				, options = that.config;
			return {
				config: options
			}
		}

		//字符常量
		, MOD_NAME = 'layTags'


		// 构造器
		, Class = function (options) {
			var that = this;
			that.config = $.extend({}, that.config, layTags.config, options);
			that.render();
		};

	//默认配置
	Class.prototype.config = {
		close: false  //默认:不开启关闭按钮
		, theme: ''   //背景:颜色
		, content: [] //默认标签
	};

	// 初始化
	Class.prototype.init = function () {
		var that = this, spans = '', options = that.config;
		$.each(options.content, function (index, item) {
			spans += '<span data-id=' + item.id + '><em>' + item.title + '</em><button type="button" class="close">×</button></span>';
			// $('<div class="layui-flow-more"><a href="javascript:;">'+ ELEM_TEXT +'</a></div>');
		});
		options.elem.before(spans);
		that.events();
	};

	Class.prototype.render = function () {
		var that = this, options = that.config;
		options.elem = $(options.elem);
	};

	Class.prototype.add = function (tagsSet) {
		if (tagsSet == null || tagsSet.length === 0) {
			return;
		}
		var that = this, spans = '', options = that.config,
			element = options.elem, existIds = [];
		var selectTags = element.parent("div.tags").children('span');
		$.each(selectTags, function () {
			existIds.push($(this).attr("data-id"));
		});
		$.each(tagsSet, function (index, item) {
			if ($.inArray(item.id.toString(), existIds) !== -1) {
				return true;
			}
			spans += '<span data-id=' + item.id + '><em>' + item.title + '</em><button type="button" class="close">×</button></span>';
			options.content.push(item);
		});
		element.before(spans);
	};

	//事件处理
	Class.prototype.events = function () {
		var that = this, options = that.config;
		options.elem.parent().on('click', '.close', function (event) {
			event.preventDefault();
			event.stopPropagation();
			let spanTag = $(this).parent('span'), id = spanTag.attr('data-id'),
				tags = spanTag.parent('div.tags');
			spanTag.remove();
			let tagSet = options.content;
			$.each(tagSet, function (index, item) {
				if (item.id == id) {
					tagSet.splice(index, 1);
					return false;
				}
			});
			if(tagSet == null || tagSet.length === 0){
				tags.children('input.lay-tags').show();
			}
		});
	};

	//核心入口
	layTags.render = function (options) {
		var inst = new Class(options);
		inst.init();
		return inst;
		//return thisLayTags.call(inst);
	};
	exports('layTags', layTags);
}).addcss('layTags.css');