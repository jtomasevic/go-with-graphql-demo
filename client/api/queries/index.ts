import { gql } from '@apollo/client';

export const GET_PIRATES = gql`
  query GetPirates {
    pirates {
      id
      name
    }
  }
`;

export const GET_CREWS = gql`
  query GetCrews {
    crews {
      id
      name
      pirates {
        id
        name
      }
    }
  }
`;

export const GET_SHIPS = gql`
  query GetShips {
    ships {
      id
      name
     }
   }
`;

export const GET_SHIP = gql`
  query GetShips{
    ships {
      id
      name
      crew {
        id
        name
        pirates{
          id
          name
        }
      }
    }
  }
`;

// export const CREATE_PIRATE = gql`
//   mutation CreatePirate($input: UpsertPirate!) {
//     createPirate(input: $input) {
//       id
//       name
//       crewId
//     }
//   }
// `;

// export const CREATE_CREW = gql`
//   mutation CreateCrew($input: UpsertCrew!) {
//     createCrew(input: $input) {
//       id
//       name
//       pirates {
//         id
//         name
//       }
//       shipId
//     }
//   }
// `;

// export const CREATE_SHIP = gql`
//   mutation CreateShip($input: UpsertShip!) {
//     createShip(input: $input) {
//       id
//       name
//       crew {
//         id
//         name
//       }
//     }
//   }
// `;