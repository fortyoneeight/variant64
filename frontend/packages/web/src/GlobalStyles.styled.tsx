import { createGlobalStyle } from 'styled-components';
export const GlobalStyles = createGlobalStyle`
    body {
        margin: 0;
        overscroll-behavior: none;
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
            'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    #root {
        min-height: 100vh;
        background-color: #336541;
    }

    h1 {
        margin: 0;
        font-size: 2em;
        color: #fff;
    }

    p {
        margin: 0;
    }

    a {
        color: #097500;
    }

    input,
    button,
    p,
    a,
    h1,
    select {
        font-family: 'texturina';
    }

    .appBody {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-between;
        min-height: 100vh;
    }

    .header {
        display: flex;
        justify-content: center;
        padding: 2vh;
    }

    .footer {
        display: flex;
        justify-content: center;
        padding: 2vh;
    }
`;
