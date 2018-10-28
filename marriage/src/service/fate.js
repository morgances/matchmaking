import { baseURL, token} from '../utils/requset';

export async function fateData(){ // 推荐人列表
	wx.request({
		url: baseURL + '/matchmaking/user/recommendusers',
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