import { connect } from 'ts-nats'
import moment from 'moment'
import request from 'request'
import es from 'event-stream'
const JSONStream = require('JSONStream')

connect({ servers: ['nats://localhost:4223'] })
  .then(nc => {
    nc.subscribe('save_article', (err, msg) => {
      console.log(msg, moment().unix())
      setTimeout(() => { msg.reply && nc.publish(msg.reply, 'ok') }, 500)
    })

    request({ url: 'http://localhost:5000/livres-final.json' })
      .pipe(JSONStream.parse('*'))
      .pipe(es.mapSync(async (data: any) => {
        nc.publish('articles', JSON.stringify(data))
      }))
})
