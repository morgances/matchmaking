import { baseURL, token} from '../utils/requset';

export async function insert(){ // 评论
	wx.request({
		url: baseURL + '/matchmaking/comment/insert',
		header:{
			"Authorization": token
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
	wx.request({
		url: baseURL + '/matchmaking/comment/ofpost',
		header:{
			"Authorization": token
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

