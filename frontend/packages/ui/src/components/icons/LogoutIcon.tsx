import { Icon } from '@chakra-ui/react';
export default function LogoutIcon(props: Parameters<typeof Icon>[0]) {
  return (
    <Icon
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="15px"
      height="14px"
      viewBox="0 0 15 14"
      fill="none"
    >
      <path
        d="M12.302 3.30771C12.3002 3.30527 12.2984 3.30287 12.2965 3.30045C12.2883 3.28962 12.2803 3.27865 12.272 3.26788L12.2712 3.26852C12.1107 3.07583 11.869 2.95312 11.5986 2.95312C11.1154 2.95312 10.7236 3.34488 10.7236 3.82813C10.7236 4.03614 10.7964 4.22708 10.9176 4.37724L10.9172 4.37756C11.4653 5.10833 11.79 6.01623 11.79 7C11.79 9.41624 9.83128 11.375 7.41504 11.375C4.99879 11.375 3.04004 9.41624 3.04004 7C3.04004 6.01623 3.3648 5.10833 3.91289 4.37756L3.91251 4.37724C4.03372 4.22708 4.10645 4.03613 4.10645 3.82813C4.10645 3.34488 3.71469 2.95312 3.23145 2.95312C2.96106 2.95312 2.71938 3.07583 2.55887 3.26852L2.55809 3.26788C2.54981 3.27864 2.54177 3.2896 2.53355 3.30043C2.53171 3.30287 2.52986 3.30529 2.52805 3.30773C1.75115 4.33439 1.29004 5.61326 1.29004 7C1.29004 10.3828 4.03229 13.125 7.41504 13.125C10.7978 13.125 13.54 10.3828 13.54 7C13.54 5.61326 13.0789 4.33438 12.302 3.30771Z"
        fill="white"
      />
      <path
        d="M7.41504 7C7.89629 7 8.29004 6.60625 8.29004 6.125V1.75C8.29004 1.26875 7.89629 0.875 7.41504 0.875C6.93379 0.875 6.54004 1.26875 6.54004 1.75V6.125C6.54004 6.60625 6.93379 7 7.41504 7Z"
        fill="white"
      />
    </Icon>
  );
}