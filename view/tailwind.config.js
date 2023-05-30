/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js}"],
  theme: {
    container:{
      center: true,
      padding: '16px'
    },
    extend: {
      colors: {
        navy: '#3c486b',
        merah: '#f45050',
        kuning: '#f9d949',
        putih: '#f0f0f0',
        second: '#ced3e2',
        twit: "#1DA1F2",
        insta: "#E4405F",
        linked: "#0A66C2",
        color2: "#D9D9D9"
      },
      fontFamily: {
        firacode: "'Fira Code', monospace;"
      },
      screens: {
        '2xl': '1320px'
      },
      keyframes: {
        'fade-in-left': {
          '0%': {
            opacity:'0',
            transform: 'translateX(-20px)'
          },
          '100%':{
            opacity: '1',
            transform: 'translateX(0)'
          },
        },

        'fade-in-rigth': {
          '0%': {
            opacity:'0',
            transform: 'translateX(20px)'
          },
          '100%':{
            opacity: '1',
            transform: 'translateX(0)'
          },
        },
      },
      animation:{
        'fade-in-left': 'fade-in-left 2s ease-in-out',
        'fade-in-rigth': 'fade-in-rigth 2s ease-in-out'
      }
    },
  },
  plugins: [],
}

