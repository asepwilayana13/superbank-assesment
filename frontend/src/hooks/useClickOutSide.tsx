import { useEffect, RefObject } from "react";

type TUseClickOutside = {
  ref: RefObject<HTMLElement | null>;
  callback: () => void;
};

const useClickOutside = ({ ref, callback }: TUseClickOutside) => {
  useEffect(() => {
    const handleClick = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        callback();
      }
    };

    document.addEventListener("click", handleClick);
    return () => {
      document.removeEventListener("click", handleClick);
    };
  }, [ref, callback]);
};

export default useClickOutside;
