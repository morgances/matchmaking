import { baseURL, token} from '../utils/requset';

export async function insert(){ // 评论
	console.log(baseURL, token,'===')
	console.log('=====service here')
	wx.request({
		url: baseURL + '/matchmaking/comment/insert',
		header:{
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
    },
    method: 'POST',
    data: {
      "user_id": '1234567890123456789012345678',
      'target_id': 1,
      'parent_id': 0,
      'content': '啦啦啦是啦啦'
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function ofPost(){ // 查看帖子的全部评论
	console.log(baseURL, token,'===')
	console.log('=====service here')
	wx.request({
		url: baseURL + '/matchmaking/comment/ofpost',
		header:{
			"Authorization": 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfaWQiOiJ0ZXN0YWNpZCIsImV4cCI6MTU0MjM0MTI2MiwiaXNfYWRtaW4iOmZhbHNlLCJvcGVuX2lkIjoiMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3OCIsInNleCI6MX0.-xCeuXQSffiYD5bbDT36cRP2gVgpwoJYYwROi1-TW9E'
    },
    method: 'POST',
    data: {
      'target_id': 1,
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

