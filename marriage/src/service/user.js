import { baseURL, token} from '../utils/requset';

export async function fillInfo(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function modifyInfo(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function changeAvatar(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function uploadPhotos(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function removePhotos(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function album(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function create(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function commend() {
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function deletePost() {
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function minePost() {
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function signIn() {
	console.log(baseURL, token,'===')
	console.log('=====service here')
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

export async function myRecord() {
	console.log(baseURL, token,'===')
	console.log('=====service here')
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