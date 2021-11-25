export default function ({ $axios, redirect }) {
    $axios.setHeader('Authorization', '456')
    $axios.onError(error => {
      if(error.response.status === 500) {
        redirect('/error')
      }
    })
  }