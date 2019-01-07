import { baseURL, setToken, token } from '../utils/requset';
import wepy from 'wepy'

export async function homeData(){ // 恋爱圈
	console.log(token, 'post')
	return await wepy.request({
    url: baseURL + '/matchmaking/post/reviewedpost',
    header:{
			"Authorization": token
		}
	})
}

export async function getDetail({ target_open_id }) {
	return await wepy.request({
		url: baseURL + '/matchmaking/user/getuserdetail',
		header:{
			"Authorization": token
		},
		method: 'POST',
    data: {
			target_open_id
    }
	})
}