import { baseURL, token} from '../utils/requset';

export async function homeData(){
	console.log(baseURL, token,'===')
	console.log('=====service here')
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