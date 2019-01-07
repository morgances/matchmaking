import { baseURL, token} from '../utils/requset';
import wepy from 'wepy'

export async function fateData(){ // 推荐人列表
	return await wepy.request({
		url: baseURL + '/matchmaking/user/recommendusers',
		header:{
			"Authorization": token
		}
	})
}