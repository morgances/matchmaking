import { baseURL, token} from '../utils/requset';

export async function fillInfo(){  // 填写个人资料
	wx.request({
		url: baseURL + '/matchmaking/user/fillinfo',
		header:{
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
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
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}