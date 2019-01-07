const baseURL = 'http://140.143.250.187:11014'
let token =  ''

export function setToken(person_token) {
  token = 'Bearer ' + person_token
  console.log(token, 'token')
  return token
}

export { baseURL, token }
