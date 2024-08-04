package sdl

//#include <SDL3/SDL_keycode.h>
import "C"

// # CategoryKeycode

// The SDL virtual key representation.
// (https://wiki.libsdl.org/SDL3/SDL_Keycode)
type Keycode C.SDL_Keycode

const K_SCANCODE_MASK = Scancode(C.SDLK_SCANCODE_MASK)

func SCANCODE_TO_KEYCODE(x Scancode) Keycode {
	return Keycode(x | K_SCANCODE_MASK)
}

const (
	K_UNKNOWN              = Keycode(C.SDLK_UNKNOWN)              /* 0 */
	K_RETURN               = Keycode(C.SDLK_RETURN)               /* '\r' */
	K_ESCAPE               = Keycode(C.SDLK_ESCAPE)               /* '\x1B' */
	K_BACKSPACE            = Keycode(C.SDLK_BACKSPACE)            /* '\b' */
	K_TAB                  = Keycode(C.SDLK_TAB)                  /* '\t' */
	K_SPACE                = Keycode(C.SDLK_SPACE)                /* ' ' */
	K_EXCLAIM              = Keycode(C.SDLK_EXCLAIM)              /* '!' */
	K_DBLAPOSTROPHE        = Keycode(C.SDLK_DBLAPOSTROPHE)        /* '"' */
	K_HASH                 = Keycode(C.SDLK_HASH)                 /* '#' */
	K_DOLLAR               = Keycode(C.SDLK_DOLLAR)               /* '$' */
	K_PERCENT              = Keycode(C.SDLK_PERCENT)              /* '%' */
	K_AMPERSAND            = Keycode(C.SDLK_AMPERSAND)            /* '&' */
	K_APOSTROPHE           = Keycode(C.SDLK_APOSTROPHE)           /* '\'' */
	K_LEFTPAREN            = Keycode(C.SDLK_LEFTPAREN)            /* '(' */
	K_RIGHTPAREN           = Keycode(C.SDLK_RIGHTPAREN)           /* ')' */
	K_ASTERISK             = Keycode(C.SDLK_ASTERISK)             /* '*' */
	K_PLUS                 = Keycode(C.SDLK_PLUS)                 /* '+' */
	K_COMMA                = Keycode(C.SDLK_COMMA)                /* ',' */
	K_MINUS                = Keycode(C.SDLK_MINUS)                /* '-' */
	K_PERIOD               = Keycode(C.SDLK_PERIOD)               /* '.' */
	K_SLASH                = Keycode(C.SDLK_SLASH)                /* '/' */
	K_0                    = Keycode(C.SDLK_0)                    /* '0' */
	K_1                    = Keycode(C.SDLK_1)                    /* '1' */
	K_2                    = Keycode(C.SDLK_2)                    /* '2' */
	K_3                    = Keycode(C.SDLK_3)                    /* '3' */
	K_4                    = Keycode(C.SDLK_4)                    /* '4' */
	K_5                    = Keycode(C.SDLK_5)                    /* '5' */
	K_6                    = Keycode(C.SDLK_6)                    /* '6' */
	K_7                    = Keycode(C.SDLK_7)                    /* '7' */
	K_8                    = Keycode(C.SDLK_8)                    /* '8' */
	K_9                    = Keycode(C.SDLK_9)                    /* '9' */
	K_COLON                = Keycode(C.SDLK_COLON)                /* ':' */
	K_SEMICOLON            = Keycode(C.SDLK_SEMICOLON)            /* ';' */
	K_LESS                 = Keycode(C.SDLK_LESS)                 /* '<' */
	K_EQUALS               = Keycode(C.SDLK_EQUALS)               /* '=' */
	K_GREATER              = Keycode(C.SDLK_GREATER)              /* '>' */
	K_QUESTION             = Keycode(C.SDLK_QUESTION)             /* '?' */
	K_AT                   = Keycode(C.SDLK_AT)                   /* '@' */
	K_LEFTBRACKET          = Keycode(C.SDLK_LEFTBRACKET)          /* '[' */
	K_BACKSLASH            = Keycode(C.SDLK_BACKSLASH)            /* '\\' */
	K_RIGHTBRACKET         = Keycode(C.SDLK_RIGHTBRACKET)         /* ']' */
	K_CARET                = Keycode(C.SDLK_CARET)                /* '^' */
	K_UNDERSCORE           = Keycode(C.SDLK_UNDERSCORE)           /* '_' */
	K_GRAVE                = Keycode(C.SDLK_GRAVE)                /* '`' */
	K_A                    = Keycode(C.SDLK_A)                    /* 'a' */
	K_B                    = Keycode(C.SDLK_B)                    /* 'b' */
	K_C                    = Keycode(C.SDLK_C)                    /* 'c' */
	K_D                    = Keycode(C.SDLK_D)                    /* 'd' */
	K_E                    = Keycode(C.SDLK_E)                    /* 'e' */
	K_F                    = Keycode(C.SDLK_F)                    /* 'f' */
	K_G                    = Keycode(C.SDLK_G)                    /* 'g' */
	K_H                    = Keycode(C.SDLK_H)                    /* 'h' */
	K_I                    = Keycode(C.SDLK_I)                    /* 'i' */
	K_J                    = Keycode(C.SDLK_J)                    /* 'j' */
	K_K                    = Keycode(C.SDLK_K)                    /* 'k' */
	K_L                    = Keycode(C.SDLK_L)                    /* 'l' */
	K_M                    = Keycode(C.SDLK_M)                    /* 'm' */
	K_N                    = Keycode(C.SDLK_N)                    /* 'n' */
	K_O                    = Keycode(C.SDLK_O)                    /* 'o' */
	K_P                    = Keycode(C.SDLK_P)                    /* 'p' */
	K_Q                    = Keycode(C.SDLK_Q)                    /* 'q' */
	K_R                    = Keycode(C.SDLK_R)                    /* 'r' */
	K_S                    = Keycode(C.SDLK_S)                    /* 's' */
	K_T                    = Keycode(C.SDLK_T)                    /* 't' */
	K_U                    = Keycode(C.SDLK_U)                    /* 'u' */
	K_V                    = Keycode(C.SDLK_V)                    /* 'v' */
	K_W                    = Keycode(C.SDLK_W)                    /* 'w' */
	K_X                    = Keycode(C.SDLK_X)                    /* 'x' */
	K_Y                    = Keycode(C.SDLK_Y)                    /* 'y' */
	K_Z                    = Keycode(C.SDLK_Z)                    /* 'z' */
	K_LEFTBRACE            = Keycode(C.SDLK_LEFTBRACE)            /* '{' */
	K_PIPE                 = Keycode(C.SDLK_PIPE)                 /* '|' */
	K_RIGHTBRACE           = Keycode(C.SDLK_RIGHTBRACE)           /* '}' */
	K_TILDE                = Keycode(C.SDLK_TILDE)                /* '~' */
	K_DELETE               = Keycode(C.SDLK_DELETE)               /* '\x7F' */
	K_PLUSMINUS            = Keycode(C.SDLK_PLUSMINUS)            /* '±' */
	K_CAPSLOCK             = Keycode(C.SDLK_CAPSLOCK)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CAPSLOCK) */
	K_F1                   = Keycode(C.SDLK_F1)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F1) */
	K_F2                   = Keycode(C.SDLK_F2)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F2) */
	K_F3                   = Keycode(C.SDLK_F3)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F3) */
	K_F4                   = Keycode(C.SDLK_F4)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F4) */
	K_F5                   = Keycode(C.SDLK_F5)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F5) */
	K_F6                   = Keycode(C.SDLK_F6)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F6) */
	K_F7                   = Keycode(C.SDLK_F7)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F7) */
	K_F8                   = Keycode(C.SDLK_F8)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F8) */
	K_F9                   = Keycode(C.SDLK_F9)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F9) */
	K_F10                  = Keycode(C.SDLK_F10)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F10) */
	K_F11                  = Keycode(C.SDLK_F11)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F11) */
	K_F12                  = Keycode(C.SDLK_F12)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F12) */
	K_PRINTSCREEN          = Keycode(C.SDLK_PRINTSCREEN)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRINTSCREEN) */
	K_SCROLLLOCK           = Keycode(C.SDLK_SCROLLLOCK)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SCROLLLOCK) */
	K_PAUSE                = Keycode(C.SDLK_PAUSE)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAUSE) */
	K_INSERT               = Keycode(C.SDLK_INSERT)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_INSERT) */
	K_HOME                 = Keycode(C.SDLK_HOME)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HOME) */
	K_PAGEUP               = Keycode(C.SDLK_PAGEUP)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEUP) */
	K_END                  = Keycode(C.SDLK_END)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_END) */
	K_PAGEDOWN             = Keycode(C.SDLK_PAGEDOWN)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PAGEDOWN) */
	K_RIGHT                = Keycode(C.SDLK_RIGHT)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RIGHT) */
	K_LEFT                 = Keycode(C.SDLK_LEFT)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LEFT) */
	K_DOWN                 = Keycode(C.SDLK_DOWN)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DOWN) */
	K_UP                   = Keycode(C.SDLK_UP)                   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UP) */
	K_NUMLOCKCLEAR         = Keycode(C.SDLK_NUMLOCKCLEAR)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_NUMLOCKCLEAR) */
	K_KP_DIVIDE            = Keycode(C.SDLK_KP_DIVIDE)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DIVIDE) */
	K_KP_MULTIPLY          = Keycode(C.SDLK_KP_MULTIPLY)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MULTIPLY) */
	K_KP_MINUS             = Keycode(C.SDLK_KP_MINUS)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MINUS) */
	K_KP_PLUS              = Keycode(C.SDLK_KP_PLUS)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUS) */
	K_KP_ENTER             = Keycode(C.SDLK_KP_ENTER)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_ENTER) */
	K_KP_1                 = Keycode(C.SDLK_KP_1)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_1) */
	K_KP_2                 = Keycode(C.SDLK_KP_2)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_2) */
	K_KP_3                 = Keycode(C.SDLK_KP_3)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_3) */
	K_KP_4                 = Keycode(C.SDLK_KP_4)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_4) */
	K_KP_5                 = Keycode(C.SDLK_KP_5)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_5) */
	K_KP_6                 = Keycode(C.SDLK_KP_6)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_6) */
	K_KP_7                 = Keycode(C.SDLK_KP_7)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_7) */
	K_KP_8                 = Keycode(C.SDLK_KP_8)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_8) */
	K_KP_9                 = Keycode(C.SDLK_KP_9)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_9) */
	K_KP_0                 = Keycode(C.SDLK_KP_0)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_0) */
	K_KP_PERIOD            = Keycode(C.SDLK_KP_PERIOD)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERIOD) */
	K_APPLICATION          = Keycode(C.SDLK_APPLICATION)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_APPLICATION) */
	K_POWER                = Keycode(C.SDLK_POWER)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_POWER) */
	K_KP_EQUALS            = Keycode(C.SDLK_KP_EQUALS)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALS) */
	K_F13                  = Keycode(C.SDLK_F13)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F13) */
	K_F14                  = Keycode(C.SDLK_F14)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F14) */
	K_F15                  = Keycode(C.SDLK_F15)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F15) */
	K_F16                  = Keycode(C.SDLK_F16)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F16) */
	K_F17                  = Keycode(C.SDLK_F17)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F17) */
	K_F18                  = Keycode(C.SDLK_F18)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F18) */
	K_F19                  = Keycode(C.SDLK_F19)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F19) */
	K_F20                  = Keycode(C.SDLK_F20)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F20) */
	K_F21                  = Keycode(C.SDLK_F21)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F21) */
	K_F22                  = Keycode(C.SDLK_F22)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F22) */
	K_F23                  = Keycode(C.SDLK_F23)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F23) */
	K_F24                  = Keycode(C.SDLK_F24)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_F24) */
	K_EXECUTE              = Keycode(C.SDLK_EXECUTE)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXECUTE) */
	K_HELP                 = Keycode(C.SDLK_HELP)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_HELP) */
	K_MENU                 = Keycode(C.SDLK_MENU)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MENU) */
	K_SELECT               = Keycode(C.SDLK_SELECT)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SELECT) */
	K_STOP                 = Keycode(C.SDLK_STOP)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_STOP) */
	K_AGAIN                = Keycode(C.SDLK_AGAIN)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AGAIN) */
	K_UNDO                 = Keycode(C.SDLK_UNDO)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_UNDO) */
	K_CUT                  = Keycode(C.SDLK_CUT)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CUT) */
	K_COPY                 = Keycode(C.SDLK_COPY)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_COPY) */
	K_PASTE                = Keycode(C.SDLK_PASTE)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PASTE) */
	K_FIND                 = Keycode(C.SDLK_FIND)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_FIND) */
	K_MUTE                 = Keycode(C.SDLK_MUTE)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MUTE) */
	K_VOLUMEUP             = Keycode(C.SDLK_VOLUMEUP)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEUP) */
	K_VOLUMEDOWN           = Keycode(C.SDLK_VOLUMEDOWN)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_VOLUMEDOWN) */
	K_KP_COMMA             = Keycode(C.SDLK_KP_COMMA)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COMMA) */
	K_KP_EQUALSAS400       = Keycode(C.SDLK_KP_EQUALSAS400)       /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EQUALSAS400) */
	K_ALTERASE             = Keycode(C.SDLK_ALTERASE)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ALTERASE) */
	K_SYSREQ               = Keycode(C.SDLK_SYSREQ)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SYSREQ) */
	K_CANCEL               = Keycode(C.SDLK_CANCEL)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CANCEL) */
	K_CLEAR                = Keycode(C.SDLK_CLEAR)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEAR) */
	K_PRIOR                = Keycode(C.SDLK_PRIOR)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_PRIOR) */
	K_RETURN2              = Keycode(C.SDLK_RETURN2)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RETURN2) */
	K_SEPARATOR            = Keycode(C.SDLK_SEPARATOR)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SEPARATOR) */
	K_OUT                  = Keycode(C.SDLK_OUT)                  /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OUT) */
	K_OPER                 = Keycode(C.SDLK_OPER)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_OPER) */
	K_CLEARAGAIN           = Keycode(C.SDLK_CLEARAGAIN)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CLEARAGAIN) */
	K_CRSEL                = Keycode(C.SDLK_CRSEL)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CRSEL) */
	K_EXSEL                = Keycode(C.SDLK_EXSEL)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_EXSEL) */
	K_KP_00                = Keycode(C.SDLK_KP_00)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_00) */
	K_KP_000               = Keycode(C.SDLK_KP_000)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_000) */
	K_THOUSANDSSEPARATOR   = Keycode(C.SDLK_THOUSANDSSEPARATOR)   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_THOUSANDSSEPARATOR) */
	K_DECIMALSEPARATOR     = Keycode(C.SDLK_DECIMALSEPARATOR)     /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_DECIMALSEPARATOR) */
	K_CURRENCYUNIT         = Keycode(C.SDLK_CURRENCYUNIT)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYUNIT) */
	K_CURRENCYSUBUNIT      = Keycode(C.SDLK_CURRENCYSUBUNIT)      /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CURRENCYSUBUNIT) */
	K_KP_LEFTPAREN         = Keycode(C.SDLK_KP_LEFTPAREN)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTPAREN) */
	K_KP_RIGHTPAREN        = Keycode(C.SDLK_KP_RIGHTPAREN)        /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTPAREN) */
	K_KP_LEFTBRACE         = Keycode(C.SDLK_KP_LEFTBRACE)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LEFTBRACE) */
	K_KP_RIGHTBRACE        = Keycode(C.SDLK_KP_RIGHTBRACE)        /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_RIGHTBRACE) */
	K_KP_TAB               = Keycode(C.SDLK_KP_TAB)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_TAB) */
	K_KP_BACKSPACE         = Keycode(C.SDLK_KP_BACKSPACE)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BACKSPACE) */
	K_KP_A                 = Keycode(C.SDLK_KP_A)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_A) */
	K_KP_B                 = Keycode(C.SDLK_KP_B)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_B) */
	K_KP_C                 = Keycode(C.SDLK_KP_C)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_C) */
	K_KP_D                 = Keycode(C.SDLK_KP_D)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_D) */
	K_KP_E                 = Keycode(C.SDLK_KP_E)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_E) */
	K_KP_F                 = Keycode(C.SDLK_KP_F)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_F) */
	K_KP_XOR               = Keycode(C.SDLK_KP_XOR)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_XOR) */
	K_KP_POWER             = Keycode(C.SDLK_KP_POWER)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_POWER) */
	K_KP_PERCENT           = Keycode(C.SDLK_KP_PERCENT)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PERCENT) */
	K_KP_LESS              = Keycode(C.SDLK_KP_LESS)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_LESS) */
	K_KP_GREATER           = Keycode(C.SDLK_KP_GREATER)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_GREATER) */
	K_KP_AMPERSAND         = Keycode(C.SDLK_KP_AMPERSAND)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AMPERSAND) */
	K_KP_DBLAMPERSAND      = Keycode(C.SDLK_KP_DBLAMPERSAND)      /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLAMPERSAND) */
	K_KP_VERTICALBAR       = Keycode(C.SDLK_KP_VERTICALBAR)       /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_VERTICALBAR) */
	K_KP_DBLVERTICALBAR    = Keycode(C.SDLK_KP_DBLVERTICALBAR)    /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DBLVERTICALBAR) */
	K_KP_COLON             = Keycode(C.SDLK_KP_COLON)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_COLON) */
	K_KP_HASH              = Keycode(C.SDLK_KP_HASH)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HASH) */
	K_KP_SPACE             = Keycode(C.SDLK_KP_SPACE)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_SPACE) */
	K_KP_AT                = Keycode(C.SDLK_KP_AT)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_AT) */
	K_KP_EXCLAM            = Keycode(C.SDLK_KP_EXCLAM)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_EXCLAM) */
	K_KP_MEMSTORE          = Keycode(C.SDLK_KP_MEMSTORE)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSTORE) */
	K_KP_MEMRECALL         = Keycode(C.SDLK_KP_MEMRECALL)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMRECALL) */
	K_KP_MEMCLEAR          = Keycode(C.SDLK_KP_MEMCLEAR)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMCLEAR) */
	K_KP_MEMADD            = Keycode(C.SDLK_KP_MEMADD)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMADD) */
	K_KP_MEMSUBTRACT       = Keycode(C.SDLK_KP_MEMSUBTRACT)       /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMSUBTRACT) */
	K_KP_MEMMULTIPLY       = Keycode(C.SDLK_KP_MEMMULTIPLY)       /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMMULTIPLY) */
	K_KP_MEMDIVIDE         = Keycode(C.SDLK_KP_MEMDIVIDE)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_MEMDIVIDE) */
	K_KP_PLUSMINUS         = Keycode(C.SDLK_KP_PLUSMINUS)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_PLUSMINUS) */
	K_KP_CLEAR             = Keycode(C.SDLK_KP_CLEAR)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEAR) */
	K_KP_CLEARENTRY        = Keycode(C.SDLK_KP_CLEARENTRY)        /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_CLEARENTRY) */
	K_KP_BINARY            = Keycode(C.SDLK_KP_BINARY)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_BINARY) */
	K_KP_OCTAL             = Keycode(C.SDLK_KP_OCTAL)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_OCTAL) */
	K_KP_DECIMAL           = Keycode(C.SDLK_KP_DECIMAL)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_DECIMAL) */
	K_KP_HEXADECIMAL       = Keycode(C.SDLK_KP_HEXADECIMAL)       /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_KP_HEXADECIMAL) */
	K_LCTRL                = Keycode(C.SDLK_LCTRL)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LCTRL) */
	K_LSHIFT               = Keycode(C.SDLK_LSHIFT)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LSHIFT) */
	K_LALT                 = Keycode(C.SDLK_LALT)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LALT) */
	K_LGUI                 = Keycode(C.SDLK_LGUI)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_LGUI) */
	K_RCTRL                = Keycode(C.SDLK_RCTRL)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RCTRL) */
	K_RSHIFT               = Keycode(C.SDLK_RSHIFT)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RSHIFT) */
	K_RALT                 = Keycode(C.SDLK_RALT)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RALT) */
	K_RGUI                 = Keycode(C.SDLK_RGUI)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_RGUI) */
	K_MODE                 = Keycode(C.SDLK_MODE)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MODE) */
	K_SLEEP                = Keycode(C.SDLK_SLEEP)                /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SLEEP) */
	K_WAKE                 = Keycode(C.SDLK_WAKE)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_WAKE) */
	K_CHANNEL_INCREMENT    = Keycode(C.SDLK_CHANNEL_INCREMENT)    /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_INCREMENT) */
	K_CHANNEL_DECREMENT    = Keycode(C.SDLK_CHANNEL_DECREMENT)    /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CHANNEL_DECREMENT) */
	K_MEDIA_PLAY           = Keycode(C.SDLK_MEDIA_PLAY)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY) */
	K_MEDIA_PAUSE          = Keycode(C.SDLK_MEDIA_PAUSE)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PAUSE) */
	K_MEDIA_RECORD         = Keycode(C.SDLK_MEDIA_RECORD)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_RECORD) */
	K_MEDIA_FAST_FORWARD   = Keycode(C.SDLK_MEDIA_FAST_FORWARD)   /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_FAST_FORWARD) */
	K_MEDIA_REWIND         = Keycode(C.SDLK_MEDIA_REWIND)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_REWIND) */
	K_MEDIA_NEXT_TRACK     = Keycode(C.SDLK_MEDIA_NEXT_TRACK)     /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_NEXT_TRACK) */
	K_MEDIA_PREVIOUS_TRACK = Keycode(C.SDLK_MEDIA_PREVIOUS_TRACK) /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PREVIOUS_TRACK) */
	K_MEDIA_STOP           = Keycode(C.SDLK_MEDIA_STOP)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_STOP) */
	K_MEDIA_EJECT          = Keycode(C.SDLK_MEDIA_EJECT)          /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_EJECT) */
	K_MEDIA_PLAY_PAUSE     = Keycode(C.SDLK_MEDIA_PLAY_PAUSE)     /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_PLAY_PAUSE) */
	K_MEDIA_SELECT         = Keycode(C.SDLK_MEDIA_SELECT)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_MEDIA_SELECT) */
	K_AC_NEW               = Keycode(C.SDLK_AC_NEW)               /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_NEW) */
	K_AC_OPEN              = Keycode(C.SDLK_AC_OPEN)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_OPEN) */
	K_AC_CLOSE             = Keycode(C.SDLK_AC_CLOSE)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_CLOSE) */
	K_AC_EXIT              = Keycode(C.SDLK_AC_EXIT)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_EXIT) */
	K_AC_SAVE              = Keycode(C.SDLK_AC_SAVE)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SAVE) */
	K_AC_PRINT             = Keycode(C.SDLK_AC_PRINT)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PRINT) */
	K_AC_PROPERTIES        = Keycode(C.SDLK_AC_PROPERTIES)        /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_PROPERTIES) */
	K_AC_SEARCH            = Keycode(C.SDLK_AC_SEARCH)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_SEARCH) */
	K_AC_HOME              = Keycode(C.SDLK_AC_HOME)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_HOME) */
	K_AC_BACK              = Keycode(C.SDLK_AC_BACK)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BACK) */
	K_AC_FORWARD           = Keycode(C.SDLK_AC_FORWARD)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_FORWARD) */
	K_AC_STOP              = Keycode(C.SDLK_AC_STOP)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_STOP) */
	K_AC_REFRESH           = Keycode(C.SDLK_AC_REFRESH)           /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_REFRESH) */
	K_AC_BOOKMARKS         = Keycode(C.SDLK_AC_BOOKMARKS)         /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_AC_BOOKMARKS) */
	K_SOFTLEFT             = Keycode(C.SDLK_SOFTLEFT)             /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTLEFT) */
	K_SOFTRIGHT            = Keycode(C.SDLK_SOFTRIGHT)            /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_SOFTRIGHT) */
	K_CALL                 = Keycode(C.SDLK_CALL)                 /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_CALL) */
	K_ENDCALL              = Keycode(C.SDLK_ENDCALL)              /* SDL_SCANCODE_TO_KEYCODE(SDL_SCANCODE_ENDCALL) */
)

// Valid key modifiers (possibly OR'd together).
// (https://wiki.libsdl.org/SDL3/SDL_Keycode)
type Keymod C.SDL_Keymod

const (
	KMOD_NONE   = Keymod(C.SDL_KMOD_NONE)   /**< no modifier is applicable. */
	KMOD_LSHIFT = Keymod(C.SDL_KMOD_LSHIFT) /**< the left Shift key is down. */
	KMOD_RSHIFT = Keymod(C.SDL_KMOD_RSHIFT) /**< the right Shift key is down. */
	KMOD_LCTRL  = Keymod(C.SDL_KMOD_LCTRL)  /**< the left Ctrl (Control) key is down. */
	KMOD_RCTRL  = Keymod(C.SDL_KMOD_RCTRL)  /**< the right Ctrl (Control) key is down. */
	KMOD_LALT   = Keymod(C.SDL_KMOD_LALT)   /**< the left Alt key is down. */
	KMOD_RALT   = Keymod(C.SDL_KMOD_RALT)   /**< the right Alt key is down. */
	KMOD_LGUI   = Keymod(C.SDL_KMOD_LGUI)   /**< the left GUI key (often the Windows key) is down. */
	KMOD_RGUI   = Keymod(C.SDL_KMOD_RGUI)   /**< the right GUI key (often the Windows key) is down. */
	KMOD_NUM    = Keymod(C.SDL_KMOD_NUM)    /**< the Num Lock key (may be located on an extended keypad) is down. */
	KMOD_CAPS   = Keymod(C.SDL_KMOD_CAPS)   /**< the Caps Lock key is down. */
	KMOD_MODE   = Keymod(C.SDL_KMOD_MODE)   /**< the !AltGr key is down. */
	KMOD_SCROLL = Keymod(C.SDL_KMOD_SCROLL) /**< the Scroll Lock key is down. */
	KMOD_CTRL   = Keymod(C.SDL_KMOD_CTRL)   /**< Any Ctrl key is down. */
	KMOD_SHIFT  = Keymod(C.SDL_KMOD_SHIFT)  /**< Any Shift key is down. */
	KMOD_ALT    = Keymod(C.SDL_KMOD_ALT)    /**< Any Alt key is down. */
	KMOD_GUI    = Keymod(C.SDL_KMOD_GUI)    /**< Any GUI key is down. */
)