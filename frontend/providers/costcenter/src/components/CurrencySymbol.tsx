import { Text, Icon, IconProps, TextProps } from '@chakra-ui/react';

export default function CurrencySymbol({
  type = 'shellCoin',
  ...props
}: {
  type?: 'shellCoin' | 'cny' | 'usd';
} & IconProps &
  TextProps) {
  return type === 'shellCoin' ? (
    <Icon w={'16px'} h="auto" fill="none" {...props} viewBox="0 0 17 16">
      <circle cx="8.5" cy="8" r="7.728" fill="#E8E8E8" stroke="#37383A" strokeWidth="0.544" />
      <circle cx="8.49995" cy="8.00002" r="6.98928" fill="#CFCFCF" />
      <path
        d="M8.50005 14.9893C12.3601 14.9893 15.4893 11.8601 15.4893 8C15.4893 6.41556 14.9621 4.95425 14.0735 3.78196C13.6263 3.71885 13.1692 3.68622 12.7045 3.68622C7.6875 3.68622 3.55904 7.49061 3.04662 12.3719C4.32758 13.9677 6.29451 14.9893 8.50005 14.9893Z"
        fill="#BEBEBE"
      />
      <circle cx="8.50008" cy="7.99998" r="5.42039" fill="#828386" />
      <path
        d="M6.26651 7.75498C6.71617 8.41242 7.64649 8.35407 7.64649 8.35407C7.41391 8.12844 7.26274 7.92226 7.24723 7.33484C7.23172 6.74743 6.89836 6.59182 6.89836 6.59182C7.49532 6.21447 7.28212 5.806 7.26274 5.35085C7.25111 5.06687 7.41779 4.8568 7.54959 4.73621C6.79101 4.85024 6.10557 5.25371 5.63615 5.8625C5.16674 6.47129 4.94958 7.23842 5.02995 8.00395C5.08422 7.85224 5.85174 7.14812 6.26651 7.75498Z"
        fill="#E8E8E8"
      />
      <path
        d="M11.7749 6.58407C11.7548 6.52013 11.7302 6.45773 11.7012 6.39734V6.39345C11.5659 6.11723 11.3416 5.89498 11.0647 5.76276C10.7877 5.63053 10.4744 5.59608 10.1755 5.66499C9.87661 5.7339 9.60967 5.90213 9.418 6.1424C9.22633 6.38266 9.12116 6.68087 9.11956 6.98865C9.11958 7.08544 9.12998 7.18196 9.15057 7.27652C9.15066 7.27781 9.15066 7.27911 9.15057 7.28041C9.15832 7.31931 9.16995 7.35821 9.18158 7.39711C9.25077 7.67116 9.26423 7.95638 9.22117 8.23575C9.1781 8.51513 9.07939 8.78294 8.93093 9.02319C8.78247 9.26344 8.5873 9.47122 8.35706 9.63411C8.12683 9.797 7.86626 9.91167 7.59089 9.97128C7.31553 10.0309 7.03102 10.0342 6.75435 9.98105C6.47768 9.92789 6.21452 9.81934 5.98057 9.66187C5.74662 9.5044 5.54669 9.30124 5.39269 9.06452C5.2387 8.8278 5.1338 8.56237 5.08427 8.28408C5.15432 8.76648 5.32453 9.22875 5.58388 9.64095C5.84324 10.0532 6.18596 10.4061 6.58986 10.6769C6.99375 10.9478 7.44983 11.1305 7.92855 11.2132C8.40727 11.296 8.89797 11.2769 9.36888 11.1572C9.83979 11.0376 10.2804 10.82 10.6622 10.5186C11.044 10.2173 11.3585 9.83875 11.5853 9.40765C11.8121 8.97655 11.9462 8.50246 11.9788 8.01606C12.0115 7.52966 11.942 7.0418 11.7749 6.58407Z"
        fill="#E8E8E8"
      />
      <path
        d="M11.1972 7.52155C11.1972 9.21671 9.8279 10.5909 8.13876 10.5909C7.24225 10.5909 6.43583 10.2038 5.87641 9.58704C5.9103 9.61299 5.94508 9.63798 5.98057 9.66187C6.21452 9.81934 6.47768 9.92789 6.75435 9.98105C7.03102 10.0342 7.31553 10.0309 7.59089 9.97128C7.86626 9.91167 8.12683 9.797 8.35706 9.63411C8.5873 9.47122 8.78247 9.26344 8.93093 9.02319C9.07939 8.78294 9.1781 8.51513 9.22117 8.23575C9.26423 7.95638 9.25077 7.67116 9.18158 7.39711C9.16995 7.35821 9.15832 7.31931 9.15057 7.28041C9.15066 7.27911 9.15066 7.27781 9.15057 7.27652C9.12998 7.18196 9.11958 7.08544 9.11956 6.98865C9.12116 6.68087 9.22633 6.38266 9.418 6.1424C9.60967 5.90213 9.87661 5.7339 10.1755 5.66499C10.2988 5.63656 10.4245 5.62568 10.5492 5.63204C10.9552 6.15298 11.1972 6.8089 11.1972 7.52155Z"
        fill="#E8E8E8"
      />
      <path
        d="M11.3335 2.79407L11.6527 3.367L12.2257 3.6862L11.6527 4.0054L11.3335 4.57833L11.0143 4.0054L10.4414 3.6862L11.0143 3.367L11.3335 2.79407Z"
        fill="#F0F0F0"
      />
    </Icon>
  ) : type === 'cny' ? (
    <Text {...props}>￥</Text>
  ) : (
    <Text {...props}>$</Text>
  );
}
