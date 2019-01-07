import { baseURL, token} from '../utils/requset';

export async function byprice(){ // 积分兑换商品列表
	return await wx.request({
		url: baseURL + '/matchmaking/goods/byprice',
		header:{
			"Authorization": token
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function tradeCreate(){ // 兑换商品
	wx.request({
		url: baseURL + '/matchmaking/trade/create',
		header:{
			"Authorization": token
    },
    method: 'POST',
    data: {
      'target_id': 1
    },
		success: function(res){
			console.log(res.data)
		},
		fail: function(err){
			console.log(err)
		}
	})
}

export async function myTrades(){ // 兑换记录
	wx.request({
		url: baseURL + '/matchmaking/trade/mytrades',
		header:{
			"Authorization": token
    },
		success: function(res){
      console.log(res)
			return res
		},
		fail: function(err){
			console.log(err)
		}
	})
}
