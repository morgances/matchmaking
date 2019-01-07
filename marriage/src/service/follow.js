import { baseURL, token} from '../utils/requset';
import wepy from 'wepy'

export async function follow(user_id){ // 700 已经关注
	return await wepy.request({
		url: baseURL + '/matchmaking/follow/follow',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      'target_open_id': user_id,
    }
	})
}

export async function unfollow(user_id){ // 取消关注
	return await wepy.request({
		url: baseURL + '/matchmaking/follow/unfollow',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      'target_open_id': user_id,
    }
	})
}

export async function following(){ // 关注的
	return await wepy.request({
		url: baseURL + '/matchmaking/follow/following',
		header:{
			"Authorization": token
    }
	})
}

export async function follower(){ // 被关注的
	return await wepy.request({
		url: baseURL + '/matchmaking/follow/follower',
		header:{
			"Authorization": token
    }
	})
}