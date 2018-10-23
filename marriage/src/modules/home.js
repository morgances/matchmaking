import { home } from '../service/home'

const Home = {
  methods: {
    async Home () {
      try {
        const resp = await home()
        if (resp) {
          return resp.data
        }
      } catch (err) {
        return false
      }
    }
  }
}

export default Home