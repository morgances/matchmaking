import { baseURL, token, setToken } from '../utils/requset';

export async function wechatLogin(code) {
	return await wx.request({
		url: baseURL + '/matchmaking/user/wechatlogin',
		method: 'POST',
		data: {
			"code": code,
		},
		success: function(res){
			console.log(res.data.data.token, 'res')
			return res.data.data.token
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function fillInfo(){  // 填写个人资料
	wx.request({
		url: baseURL + '/matchmaking/user/fillinfo',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      "phone": '12222222222',
      "wechat": 'wbfwdd',
      "sex": 1,
      "real_name": 'Wbofeng',
      "birthday": '1998-04-04',
      "height": '180',
      "job": '程序员',
      "faith": '无',
      "self_introduction": '是人',
      "selec_criteria": '是人就行'
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function modifyInfo(){ // 修改个人资料
	wx.request({
		url: baseURL + '/matchmaking/user/modifyinfo',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      "phone": '12222222222',
      "wechat": 'wbfwdd',
      "faith": '无',
      "self_introduction": '是人',
			"selec_criteria": '是人就行',
			"nick_name": '啦啦啦啦'
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function changeAvatar(){ // 更换头像
	wx.request({
		url: baseURL + '/matchmaking/user/changeavatar',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {},
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function uploadPhotos(){ // 上传照片
	wx.request({
		url: baseURL + '/matchmaking/user/uploadphotos',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {},
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function removePhotos(){ // 删除照片
	wx.request({
		url: baseURL + '/matchmaking/user/removephotos',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {},
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function album(){ // 获取相册
	wx.request({
		url: baseURL + '/matchmaking/user/album',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      "target_open_id": '1234567890123456789012345678'
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function create(){ // 发表动态
	wx.request({
		url: baseURL + '/matchmaking/post/create',
		header:{
			"Authorization": token
		},
		method: 'POST',
    data: {
			"title": 'eee',
			"content": 'sss',
			"image_num": 0
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function commend() { // 点赞
	wx.request({
		url: baseURL + '/matchmaking/post/commend',
		header:{
			"Authorization": token
		},
		method: 'POST',
    data: {
			"target_id": 1
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function deletePost() { // 删除动态
	wx.request({
		url: baseURL + '/matchmaking/post/userdelete',
		header:{
			"Authorization": token
		},
		method: 'POST',
    data: {
			"target_id": 1
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function minePost() { // 我的动态
	wx.request({
		url: baseURL + '/matchmaking/post/mine',
		header:{
			"Authorization": token
		},
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function signIn() { // 签到
	wx.request({
		url: baseURL + '/matchmaking/signin/signin',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {},
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function myRecord() { //签到记录
	wx.request({
		url: baseURL + '/matchmaking/signin/myrecord',
		header:{
			"Authorization": token
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}