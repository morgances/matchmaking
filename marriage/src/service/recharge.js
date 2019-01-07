import { baseURL, token } from '../utils/requset';

export async function rechargeVip(){ // 兑换 Vip
	wx.request({
		url: baseURL + '/matchmaking/recharge/vip',
		header:{
			"Authorization": token
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

export async function rechargeRose(){ // 兑换 玫瑰花
	console.log(token, 'token')
	wx.request({
		url: baseURL + '/matchmaking/recharge/rose',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      'rose_num': 1
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}