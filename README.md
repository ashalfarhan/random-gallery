<h1 align="center">Gallery API</h1>

<hr>

A simple REST API for mocking your Instagram/Unsplash/Pinterest clone project. This API just response the randomized json data. And the image source is actually from [Picsum Photos](https://picsum.photos)

## Endpoints

- `/api/images`

Get list of images

### Query Params

- limit
  - type: `number`
  - default: `5`
  - note: must be greater than `0`, and less than `25`

- page
  - type: `number`
  - default: `1`
