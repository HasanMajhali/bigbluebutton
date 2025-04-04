import styled, { css } from 'styled-components';
import { fontSizeLarge, fontSizeBase } from '/imports/ui/stylesheets/styled-components/typography';
import {
  smPaddingX,
  borderSize,
  borderSizeLarge,
  borderSizeSmall,
} from '/imports/ui/stylesheets/styled-components/general';
import {
  colorWhite,
  colorPrimary,
  colorDanger,
  colorDangerDark,
  colorGray,
  btnDefaultGhostBg,
} from '/imports/ui/stylesheets/styled-components/palette';
import SpinnerStyles from '/imports/ui/components/common/loading-screen/styles';

interface RecordingIndicatorIconProps {
  titleMargin: boolean;
}

interface RecordingIndicatorProps {
  recording: boolean;
  disabled: boolean;
  time: number;
  isPhone?: boolean;
}

interface RecordingStatusViewOnlyProps {
  recording: boolean;
}

interface SpinnerOverlayProps {
  animations: boolean;
}

const RecordingIndicatorIcon = styled.span<RecordingIndicatorIconProps>`
  width: ${fontSizeLarge};
  height: ${fontSizeLarge};
  font-size: ${fontSizeBase};
  user-select: none;

  ${({ titleMargin }) => titleMargin && `
    [dir="ltr"] & {
      margin-right: ${smPaddingX};
    }
  `}
`;

const RecordingControl = styled.button<RecordingIndicatorProps>`
  display: flex;
  align-items: center;
  user-select: none;
  background: ${btnDefaultGhostBg};

  span {
    border: none;
    box-shadow: none;
    background-color: transparent !important;
    color: ${colorWhite} !important;
  }

  &:hover:not(:disabled) {
    color: ${colorWhite} !important;
    cursor: pointer;
  }

  &:focus {
    outline: none;
    box-shadow: 0 0 0 ${borderSize} ${colorPrimary};
  }

  ${({ recording }) => recording && `
    padding: 0.5rem 1rem;
    background-color: ${colorDanger};
    border: ${borderSizeSmall}} solid ${colorDanger};
    border-radius: 2em 2em;

    &:focus {
      border: ${borderSizeLarge};
      box-shadow: none;
    }
  `}

  ${({ recording }) => !recording && `
    padding: 0.5rem 1rem;
    border: ${borderSizeSmall};
    border-radius: 2em 2em;

    &:focus {
      border: ${borderSizeLarge};
      box-shadow: none;
    }
  `}

  ${({ disabled, time }) => disabled && time <= 0 && css`
    display: none;
  `}

  ${({ disabled, time }) => disabled && time > 0 && css`
    cursor: not-allowed;
    pointer-events: none;
  `}
`;

const PresentationTitle = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  font-weight: 400;
  color: ${colorWhite};
  font-size: ${fontSizeBase};
  padding: 0;
  margin-right: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  max-width: 30vw;

  [dir='rtl'] & {
    margin-left: 0;
    margin-right: ${smPaddingX};
  }

  & > [class^='icon-bbb-'] {
    font-size: 75%;
  }

  span {
    vertical-align: middle;
  }
`;

const VisuallyHidden = styled.span`
  position: absolute;
  overflow: hidden;
  clip: rect(0 0 0 0);
  height: 1px;
  width: 1px;
  margin: -1px;
  padding: 0;
  border: 0;
`;

const PresentationTitleSeparator = styled.span`
  color: ${colorGray};
  font-size: ${fontSizeBase};
  margin: 0 1rem;
`;

const RecordingIndicator = styled.div<RecordingIndicatorProps>`
  padding-right: 1rem;

  ${({ isPhone }) => isPhone && `
    margin-left: ${smPaddingX};
  `}

  &:hover {
    outline: transparent;
    outline-style: dotted;
    outline-width: ${borderSize};
  }

  &:active,
  &:focus,
  &:focus-within {
    outline: transparent;
    outline-width: ${borderSize};
    outline-style: solid;
  }
`;

const RecordingStatusViewOnly = styled.div<RecordingStatusViewOnlyProps>`
  display: flex;

  ${({ recording }) => recording && `
    padding: 5px 5px 5px 5px;
    background-color: ${colorDangerDark};
    border: ${borderSizeLarge} solid ${colorDangerDark};
    border-radius: 10px;
  `}
`;

const SpinnerOverlay = styled(SpinnerStyles.Spinner)<SpinnerOverlayProps>`
  & > div {
    background-color: ${colorWhite};;
    height: 0.5625rem;
    width: 0.5625rem;
  }
`;

const Bounce1 = styled(SpinnerStyles.Bounce1)<SpinnerOverlayProps>`
  height: 0.5625rem;
  width: 0.5625rem;
`;

const Bounce2 = styled(SpinnerStyles.Bounce2)<SpinnerOverlayProps>`
  height: 0.5625rem;
  width: 0.5625rem;
`;

export default {
  RecordingIndicatorIcon,
  RecordingControl,
  PresentationTitle,
  VisuallyHidden,
  PresentationTitleSeparator,
  RecordingIndicator,
  RecordingStatusViewOnly,
  SpinnerOverlay,
  Bounce1,
  Bounce2,
};
