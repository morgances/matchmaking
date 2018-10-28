import { baseURL } from '../utils/requset';

export async function homeData(){ // 恋爱圈
	wx.request({
		url: baseURL + '/matchmaking/post/many?isreviewed=true',
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}