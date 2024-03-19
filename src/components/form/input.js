import { forwardRef } from "react";

const Input = forwardRef((props, ref) => {
  return (
    <div className="mb-3">
      <label htmlFor="{props.name}" className="form-label">
        {props.title}
      </label>
      <input
        type={props.type}
        className={props.className}
        id={props.name}
        placeholder={props.placeholder}
        onChange={props.onChange}
        autocomplete={props.autocomplete}
        value={props.value}
      />
      <div className={props.errorDiv}>{props.ErrorMsg}</div>
    </div>
  );
});

export default Input;
