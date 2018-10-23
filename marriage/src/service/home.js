import request from '../utils/requset';

export async function homeData() {
  try {
		const resp = await request({
			url: '',
			method: 'get'
    })
		return resp
	} catch (err) {
		console.log(err)
		return false
	}
}