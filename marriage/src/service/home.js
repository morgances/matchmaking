import { baseURL, token} from '../utils/requset';

export async function homeData(){
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