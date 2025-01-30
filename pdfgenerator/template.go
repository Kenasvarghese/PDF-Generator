package pdfgenerator

const SampleTemplate = `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Personal Details</title>
  </head>
  <body>
    <div class="main-div">
      <div class="heading-div">
        <!--<img
          class="img"
          alt="profile"
          src="https://www.google.com/url?sa=i&url=https%3A%2F%2Fimagekit.io%2Fblog%2Fwhat-is-image-cdn-guide%2F&psig=AOvVaw2s72H6aveUrcKeOAP0WHdZ&ust=1737025189259000&source=images&cd=vfe&opi=89978449&ved=0CBEQjRxqFwoTCNil0aHJ94oDFQAAAAAdAAAAABAE"
        /> -->
        <h2 class="name-heading">{{name}}</h2>
        <p class="email-heading">{{email}}</p>
      </div>
      <div class="details-div">
        <span class="details-header">Personal Details</span>
        <span class="details-key">Full name:</span>
        <span class="details-value">{{name}}</span>
        <span class="details-key">Date of birth:</span>
        <span class="details-value">{{dob}}</span
        ><span class="details-key">Gender:</span>
        <span class="details-value">{{gender}}</span
        ><span class="details-key rounded-bl">E-mail:</span>
        <span class="details-value rounded-br">{{email}}</span>
      </div>
    </div>
  </body>
  <style>
    .rounded-bl {
      border-radius: 0 0 0 1rem;
    }
    .rounded-br {
      border-radius: 0 0 1rem 0;
    }
    .details-header {
      border: 1.5px solid rgb(238, 238, 238);
      border-radius: 1rem 1rem 0 0;
      padding: 1rem;
      font-size: 0.9rem;
      grid-column: span 2;
      font-weight: 600;
    }
    .details-key {
      border: 1.5px solid rgb(233, 231, 231);
      border-top: 0;
      white-space: nowrap;
      border-right: 0;
      padding: 1rem;
      color: gray;
      font-weight: 600;
      font-size: 0.8rem;
    }
    .details-value {
      border: 1.5px solid rgb(233, 231, 231);
      border-top: 0;
      border-left: 0;
      padding: 1rem;
      text-align: left;
      font-weight: 600;
      font-size: 0.8rem;
    }
    .details-div {
      min-width: 300px;
      display: grid;
      grid-template-columns: 0.4fr 0.6fr;
    }
    body {
      background-color: rgb(248, 248, 248);
      min-height: 100vh;
      font-family: Verdana, Geneva, Tahoma, sans-serif;
      margin: 0;
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
    }
    .main-div {
      max-width: 40rem;
      display: flex;
      background-color: white;
      flex-direction: column;
      justify-content: center;
      gap: 0.5rem;
      width: fit-content;
      padding: 2rem;
      border: 1px solid rgb(230, 230, 230);
      border-radius: 1rem;
      box-shadow: 2px 2px 10px rgb(238, 238, 238);
    }
    .heading-div {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      gap: 0.25rem;
      padding-block: 1rem;
    }
    .img {
      height: 3.5rem;
      width: 3.5rem;
      margin-bottom: 0.5rem;
      margin-inline: auto;
      border: rgb(189, 185, 185) 1px solid;
      border-radius: 100%;
    }
    .name-heading {
      padding: 0;
      margin: 0;
      font-size: 1.5rem;
    }
    .email-heading {
      margin: 0;
      font-size: small;
      color: gray;
    }
  </style>
</html>
`
