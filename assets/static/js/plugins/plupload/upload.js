function OSSUpload(up_locale) {
	var accessid = '', host = '', policyBase64 = '', signature = '',
	callbackbody = '', filePath = up_locale.filePath, expire = 0;
	var now = timestamp = Date.parse(new Date()) / 1000;
	function send_request() {
		var location = window.location;
		var basePath = location.protocol +"//"+ location.host+ctx;
		var local = 1;
		var xmlhttp = null;
		if (window.XMLHttpRequest) {
			xmlhttp = new XMLHttpRequest();
		} else if (window.ActiveXObject) {
			xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
		}
		if (xmlhttp != null) {// serverUrl是 用户获取 '签名和Policy' 等信息的应用服务器的URL，请将下面的IP和Port配置为您自己的真实信息。
			var p = "dir="+encodeURI(encodeURI(filePath))+"&local="+local;
			var serverUrl = basePath+'OSSCallback?'+p;
			xmlhttp.open("GET", serverUrl, false);
			xmlhttp.send(null);
			return xmlhttp.responseText
		} else {
			$.modal.msgError("您的浏览器不支持 XMLHTTP 请求！");
		}
	}

	function get_signature() {// 可以判断当前expire是否超过了当前时间， 如果超过了当前时间， 就重新取一下，3s 作为缓冲。
		now = timestamp = Date.parse(new Date()) / 1000;
		if (expire < now + 3) {
			body = send_request();
			var obj = eval("(" + body + ")");
			host = obj['host'];
			policyBase64 = obj['policy'];
			accessid = obj['accessid'];
			signature = obj['signature'];
			expire = parseInt(obj['expire']);
			callbackbody = obj['callback'];
			return true;
		}
		return false;
	}

	function set_upload_param(up, filename) {
		var pos = filename.lastIndexOf('.'), suffix = '';
		if (pos !== -1) {
			suffix = filename.substring(pos).toLowerCase();
		}
		var g_object_name = filePath+up_locale.id+suffix;
		new_multipart_params = {
			'key': g_object_name,
			'policy': policyBase64,
			'OSSAccessKeyId': accessid,
			'success_action_status': '200', //让服务端返回200,不然，默认会返回204
			'callback': callbackbody,
			'signature': signature
		};
		up.setOption({
			'url': host,
			'multipart_params': new_multipart_params
		});
	}

	function handel_error(up) {
		up.stop();
		layer.closeAll('loading');
		$.modal.msgError(up_locale.failed);
		up.destroy();
	}

	this.start_upload = function () {
		get_signature();//获取签名
		this.uploader.start();
	};

	this.uploader = new plupload.Uploader({
		runtimes: 'html5,flash,silverlight,html4',
		browse_button: 'selectFiles',
		multi_selection: false,
		container: document.getElementById('container'),
		flash_swf_url: 'Moxie.swf',
		silverlight_xap_url: 'Moxie.xap',
		url: 'http://oss.aliyuncs.com',
		filters: {
			mime_types: [
				{ title : "files", extensions : "mp3,mp4,pdf,docx"}
			],
			max_file_size: up_locale.fileSize, //最大能上传的文件大小
			prevent_duplicates: true //不允许选取重复文件
		},

		init: {
			PostInit: function () {
				$('#filesInfo').empty();
			},

			FilesAdded: function (up, files) {
				//最多上传的文件数
				if(up.files.length > up_locale.fileNum){
					/*$.modal.alertWarning(up_locale.maxNum);
					up.splice(max_num, 999);*/
					up.splice(0, 1);//删除第一个文件
				}
				var filesInfo = $('#filesInfo').empty();
				plupload.each(up.files, function (file) {
					filesInfo.append('<div id="' + file.id + '">' + file.name + ' (' + plupload.formatSize(file.size) + ')<b></b>'
						+ '<div class="progress"><div class="progress-bar" style="width: 0%"></div></div>'
						+ '</div>');
				});
			},

			BeforeUpload: function (up, file) {
				set_upload_param(up, file.name);
			},

			UploadProgress: function (up, file) {
				var d = document.getElementById(file.id);
				d.getElementsByTagName('b')[0].innerHTML = '<span>' + file.percent + "%</span>";
				var prog = d.getElementsByTagName('div')[0];
				var progBar = prog.getElementsByTagName('div')[0];
				progBar.style.width = 4 * file.percent + 'px';
				progBar.setAttribute('aria-valuenow', file.percent);
			},

			FileUploaded: function (up, file, info) {
				if (info.status === 200 || info.status === 203) {
					layer.msg(up_locale.success, {icon:1, time:2000, shift:5});
					console.log(info.status===200 ? ("Object Name："+file.name) : "OSS Callback Failed");
					console.log(info.response);
					setTimeout(function () {
						window.location.reload();
					}, 2000);
				}else {
					handel_error(up);
					console.log('其它原因：'+info.response);
				}
			},

			/*UploadComplete:function(up, file){
				window.location.href="/success";
			},*/

			Error: function (up, err) {
				if (err.code === -600) {
					$.modal.msgError(up_locale.maxSize);
				}else if (err.code === -601) {
					$.modal.msgError(up_locale.support);
				}else if (err.code === -602) {
					$.modal.msgError(up_locale.repeat);
				}else {
					handel_error(up);
					console.log(err.response);
				}
			},

			Destroy:function(up){
				createUploader();
			}
		}
	});
	this.uploader.init();
	return this;
}
