package email

// TODO: Vérifier si c'est un soft ou hard bounce => informer l'utilisateur
// TODO: Ajouter le deuxième SMTP
// TODO: Ajouter ScheduledAt dans la fonction sendMail()

import (
    "context"
    "errors"
    "fmt"
    "log"
    "reflect"

    "github.com/cbroglie/mustache"
    brevo "github.com/sendinblue/APIv3-go-library/lib"
)

type APIClient struct {
    Context context.Context
    Sib     *brevo.APIClient
}

const BREVO_API_KEY = "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"

func SendInit() APIClient {
    var context context.Context
    cfg := brevo.NewConfiguration()
    //Configure API key authorization: api-key
    cfg.AddDefaultHeader("api-key", BREVO_API_KEY)
    //Configure API key authorization: partner-key
    api_client := APIClient{}
    api_client.Sib = brevo.NewAPIClient(cfg)
    api_client.Context = context
    return api_client
}

func SetToEmail(to []any) []brevo.SendSmtpEmailTo {
    var tos []brevo.SendSmtpEmailTo
    string_email := make([]string, len(to))
    for i, t := range to {
        string_email[i] = fmt.Sprint(t)
        to_smtp := brevo.SendSmtpEmailTo{
            Email: string_email[i],
        }
        tos = append(tos, to_smtp)
    }

    return tos
}

func SetTemplate(template string, variable map[string]any) (string, error) {

    data, err := mustache.Render(template, variable)
    if err != nil {
        return "", err
    }

    return data, nil
}

func SetAttach(local_attachs []any) []brevo.SendSmtpEmailAttachment {
    var attachs []brevo.SendSmtpEmailAttachment
    for _, attach := range local_attachs {
        /*picture_b, err := builtins.GetPBR(attach.(map[string]any)["item"].(string))
        if err != nil {
            continue
        }*/

        //sEnc := b64.StdEncoding.EncodeToString([]byte(picture_b))
        attach_t := brevo.SendSmtpEmailAttachment{
            Url:  attach.(map[string]any)["item"].(string),
            Name: attach.(map[string]any)["name"].(string),
        }
        attachs = append(attachs, attach_t)
    }

    return attachs
}

func SendMail(data map[string]any) (bool, error) {

    // set settings instance
    api_client := SendInit()

    // set sendSmtpEmail configuration
    sender := brevo.SendSmtpEmailSender{}
    sender.Email = data["from_email"].(string)
    sender.Name = data["from"].(string)
    to_email := SetToEmail(data["to"].([]any))
    var template string
    var err error
    if data["replace"] != nil {
        template, err = SetTemplate(data["template"].(string), data["replace"].(map[string]any))
        if err != nil {
            return false, err
        }
    } else {
        template = data["template"].(string)
    }

    var attachments []brevo.SendSmtpEmailAttachment
    if data["attach"] != nil {
        log.Println(reflect.TypeOf(data["attach"]))
        attachments = SetAttach(data["attach"].([]any))
    } else {
        attachments = nil
    }

    // send email
    _, http_code, err := api_client.Sib.TransactionalEmailsApi.SendTransacEmail(api_client.Context, brevo.SendSmtpEmail{
        Sender:      &sender,
        To:          to_email,
        HtmlContent: template,
        Subject:     data["subject"].(string),
        Attachment:  attachments,
    })

    if err != nil {
        return false, err
    }

    if (http_code.StatusCode != 200) && (http_code.StatusCode != 201) {
        return false, errors.New("http error")
    }

    return true, nil
}
